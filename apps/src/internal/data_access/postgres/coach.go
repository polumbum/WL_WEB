package dataaccess

import (
	"errors"
	"log"
	"src/internal/converters"
	dataaccess "src/internal/data_access"
	"src/internal/data_access/models"
	"src/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CoachRepository struct {
	db *gorm.DB
}

func NewCoachRepository(db *gorm.DB) *CoachRepository {
	return &CoachRepository{
		db: db,
	}
}

func (r *CoachRepository) Update(coach *domain.Coach) error {
	model, err := converters.NewCoachConverter().ToModel(coach)
	if err != nil {
		return err
	}

	err = r.db.Save(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dataaccess.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *CoachRepository) Create(coach *domain.Coach) error {
	model, err := converters.NewCoachConverter().ToModel(coach)
	if err != nil {
		return err
	}

	model.ID = uuid.New()
	err = r.db.Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CoachRepository) ListCoaches(
	page int,
	batch int,
	sortStr string,
	filterStr string,
) ([]*domain.Coach, error) {
	if page < 1 {
		page = dataaccess.DefaultPage
	}
	if batch < 1 {
		batch = dataaccess.DefaultBatch
	}
	coaches := []*models.Coach{}
	offset := (page - 1) * batch

	filter := NewCoachFilter(filterStr)
	query := r.db.Model(&models.Coach{})
	query = filter.Apply(query)

	err := query.Order(sortStr).
		Offset(offset).
		Limit(batch).
		Find(&coaches).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	resDomain := []*domain.Coach{}
	conv := converters.NewCoachConverter()
	for _, item := range coaches {
		c, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		resDomain = append(resDomain, c)
	}
	return resDomain, nil
}

func (r *CoachRepository) GetCoachByID(coachID uuid.UUID) (
	*domain.Coach,
	error,
) {
	coach := &models.Coach{}
	err := r.db.Where("id = ?", coachID).First(coach).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}
	domain, err := converters.NewCoachConverter().ToDomain(coach)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

func (r *CoachRepository) ListSportsmen(
	coachID uuid.UUID,
	page int,
	batch int,
	sortStr string,
	filterStr string,
) (
	[]*domain.Sportsman,
	error,
) {
	if page < 1 {
		page = dataaccess.DefaultPage
	}
	if batch < 1 {
		batch = dataaccess.DefaultBatch
	}

	sportsmen := []*models.Sportsman{}
	offset := (page - 1) * batch

	filter := NewSportsmanFilter(filterStr)
	query := r.db.Table("sportsmen").
		Joins("INNER JOIN sportsmen_coaches ON sportsmen.id = sportsmen_coaches.sportsman_id").
		Joins("INNER JOIN coaches ON coaches.id = sportsmen_coaches.coach_id AND coaches.id = ?",
			coachID)
	query = filter.Apply(query)

	/*err := r.db.Table("sportsmen").
	Joins("INNER JOIN sportsmen_coaches ON sportsmen.id = sportsmen_coaches.sportsman_id").
	Joins("INNER JOIN coaches ON coaches.id = sportsmen_coaches.coach_id AND coaches.id = ?",
		coachID).
	Find(&sportsmen).Error*/
	err := query.Order(sortStr).Offset(offset).Limit(batch).Find(&sportsmen).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	resDomain := []*domain.Sportsman{}
	conv := converters.NewSportsmanConverter()
	for _, item := range sportsmen {
		s, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		resDomain = append(resDomain, s)
	}

	return resDomain, nil
}

func (r *CoachRepository) AddSportsman(coachID, smID uuid.UUID) (
	*domain.SportsmenCoach,
	error,
) {
	rec := &models.SportsmenCoach{
		SportsmanID: smID,
		CoachID:     coachID,
	}
	err := r.db.Create(rec).Error
	if err != nil {
		return nil, err
	}
	domain, err := converters.NewSmCoachesConverter().
		ToDomain(rec)
	if err != nil {
		return nil, err
	}
	return domain, nil
}

func (r *CoachRepository) RemoveSportsman(coachID, smID uuid.UUID) error {
	var smCoach models.SportsmenCoach
	result := r.db.
		Where("coach_id = ? AND sportsman_id = ?", coachID, smID).
		Delete(&smCoach)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return dataaccess.ErrNotFound
	}

	return nil
}

func (r *CoachRepository) Delete(cID uuid.UUID) error {
	err := r.db.Delete(&models.Coach{}, cID).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
