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

type CompetitionRepository struct {
	db *gorm.DB
}

func NewCompetitionRepository(db *gorm.DB) *CompetitionRepository {
	return &CompetitionRepository{
		db: db,
	}
}

func (r *CompetitionRepository) Update(
	competition *domain.Competition,
) error {
	model, err := converters.NewCompConverter().
		ToModel(competition)
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

func (r *CompetitionRepository) Create(
	competition *domain.Competition,
) error {
	model, err := converters.NewCompConverter().
		ToModel(competition)
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

func (r *CompetitionRepository) GetCompetitionByID(
	competitionID uuid.UUID,
) (*domain.Competition,
	error,
) {

	competition := &models.Competition{}
	err := r.db.Where("id = ?", competitionID).First(competition).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain, err := converters.NewCompConverter().ToDomain(competition)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *CompetitionRepository) ListCompetitions(
	page int,
	batch int,
	sortStr string,
	filterStr string,
) (
	[]*domain.Competition,
	error,
) {
	if page < 1 {
		page = dataaccess.DefaultPage
	}
	if batch < 1 {
		batch = dataaccess.DefaultBatch
	}

	competitions := []*models.Competition{}
	offset := (page - 1) * batch

	filter := NewCompFilter(filterStr)
	query := r.db.Model(&models.Competition{})
	query = filter.Apply(query)

	//err := r.db.Find(&competitions).Error
	err := query.Order(sortStr).
		Offset(offset).
		Limit(batch).
		Find(&competitions).
		Error
	if err != nil {
		return nil, err
	}

	domain := []*domain.Competition{}
	conv := converters.NewCompConverter()

	for _, item := range competitions {
		c, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, c)
	}
	return domain, nil
}

func (r *CompetitionRepository) ListByOrgID(id uuid.UUID) (
	[]*domain.Competition,
	error,
) {
	comps := []*models.Competition{}

	err := r.db.Where("org_id = ?", id).
		Find(&comps).Error
	if err != nil {
		return nil, err
	}

	domain := []*domain.Competition{}
	conv := converters.NewCompConverter()
	for _, item := range comps {
		r, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, r)
	}
	return domain, nil
}

func (r *CompetitionRepository) RegisterSportsman(
	compApplication *domain.CompApplication,
) error {
	model, err := converters.NewCompApplConverter().
		ToModel(compApplication)
	if err != nil {
		return err
	}

	err = r.db.Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CompetitionRepository) DeleteRegistration(
	smID,
	compID uuid.UUID,
) error {
	err := r.db.Where("sportsman_id = ? AND competition_id = ?", smID, compID).Delete(&models.CompApplication{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dataaccess.ErrNotFound
		}
	}
	return nil
}

func (r *CompetitionRepository) GetUpcoming(smID uuid.UUID) (
	[]*domain.Competition,
	error,
) {
	comps := []*models.Competition{}
	if err := r.db.Raw("SELECT * FROM get_upcoming_comps_sm(?)", smID).Find(&comps).Error; err != nil {
		return nil, err
	}

	domain := []*domain.Competition{}
	conv := converters.NewCompConverter()
	for _, item := range comps {
		c, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, c)
	}

	return domain, nil
}
func (r *CompetitionRepository) ListUpcoming() (
	[]*domain.Competition,
	error,
) {
	comps := []*models.Competition{}
	if err := r.db.Raw("SELECT * FROM get_upcoming_comps()").Find(&comps).Error; err != nil {
		return nil, err
	}

	domain := []*domain.Competition{}
	conv := converters.NewCompConverter()
	for _, item := range comps {
		c, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, c)
	}

	return domain, nil
}

func (r *CompetitionRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&models.Competition{}, id).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
