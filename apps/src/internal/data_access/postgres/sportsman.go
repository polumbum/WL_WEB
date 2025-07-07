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

type SportsmanRepository struct {
	db *gorm.DB
}

func NewSportsmanRepository(db *gorm.DB) *SportsmanRepository {
	return &SportsmanRepository{
		db: db,
	}
}

func (r *SportsmanRepository) Update(sportsman *domain.Sportsman) (
	*domain.Sportsman,
	error,
) {
	model, err := converters.NewSportsmanConverter().ToModel(sportsman)
	if err != nil {
		return nil, err
	}

	err = r.db.Save(model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}
	return sportsman, nil
}

func (r *SportsmanRepository) Create(sportsman *domain.Sportsman) error {
	model, err := converters.NewSportsmanConverter().ToModel(sportsman)
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

func (r *SportsmanRepository) ListSportsmen(
	page int,
	batch int,
	sortStr string,
	filterStr string,
) ([]*domain.Sportsman, error) {
	if page < 1 {
		page = dataaccess.DefaultPage
	}
	if batch < 1 {
		batch = dataaccess.DefaultBatch
	}

	sportsmen := []*models.Sportsman{}
	offset := (page - 1) * batch

	filter := NewSportsmanFilter(filterStr)
	query := r.db.Model(&models.Sportsman{})
	query = filter.Apply(query)

	err := query.Order(sortStr).Offset(offset).Limit(batch).Find(&sportsmen).Error

	//err := r.db.Order(sortStr).Offset(offset).Limit(batch).Find(&sportsmen).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain := []*domain.Sportsman{}
	conv := converters.NewSportsmanConverter()
	for _, item := range sportsmen {
		c, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, c)
	}
	return domain, nil
}

func (r *SportsmanRepository) GetSportsmanByID(sportsmanID uuid.UUID) (
	*domain.Sportsman,
	error,
) {
	sportsman := &models.Sportsman{}
	err := r.db.Where("id = ?", sportsmanID).First(sportsman).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain, err := converters.NewSportsmanConverter().ToDomain(sportsman)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *SportsmanRepository) ListResults(sportsmanID uuid.UUID) (
	[]*domain.Result,
	error,
) {
	results := []*models.Result{}
	err := r.db.Where("sportsman_id = ?", sportsmanID).Find(&results).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain := []*domain.Result{}
	conv := converters.NewResultConverter()
	for _, item := range results {
		r, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, r)
	}
	return domain, nil
}

func (r *SportsmanRepository) Delete(smID uuid.UUID) error {
	/*err := r.db.Delete(&domain.SportsmenCoach{}, smID).Error
	if err != nil {
		log.Println(err)
		return err
	}*/
	err := r.db.Delete(&models.Sportsman{}, smID).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
