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

type TCampRepository struct {
	db *gorm.DB
}

func NewTCampRepository(db *gorm.DB) *TCampRepository {
	return &TCampRepository{
		db: db,
	}
}

func (r *TCampRepository) Update(tCamp *domain.TCamp) error {
	model, err := converters.NewTCampConverter().ToModel(tCamp)
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

func (r *TCampRepository) Create(tCamp *domain.TCamp) error {
	model, err := converters.NewTCampConverter().ToModel(tCamp)
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

func (r *TCampRepository) GetTCampByID(tCampID uuid.UUID) (
	*domain.TCamp,
	error,
) {
	tCamp := &models.TCamp{}
	err := r.db.Where("id = ?", tCampID).First(tCamp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, dataaccess.ErrNotFound
		}
		return nil, err
	}

	domain, err := converters.NewTCampConverter().ToDomain(tCamp)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *TCampRepository) ListTCamps(
	page int,
	batch int,
	sortStr string,
	filterStr string,
) ([]*domain.TCamp, error) {
	if page < 1 {
		page = dataaccess.DefaultPage
	}
	if batch < 1 {
		batch = dataaccess.DefaultBatch
	}

	tCamps := []*models.TCamp{}
	offset := (page - 1) * batch

	//err := r.db.Find(&tCamps).Error
	filter := NewTCampFilter(filterStr)
	log.Println(filter)
	query := r.db.Model(&models.TCamp{})
	query = filter.Apply(query)

	err := query.Order(sortStr).
		Offset(offset).
		Limit(batch).
		Find(&tCamps).
		Error
	if err != nil {
		return nil, err
	}

	domain := []*domain.TCamp{}
	conv := converters.NewTCampConverter()
	for _, item := range tCamps {
		r, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, r)
	}
	return domain, nil
}

func (r *TCampRepository) ListByOrgID(id uuid.UUID) (
	[]*domain.TCamp,
	error,
) {
	tCamps := []*models.TCamp{}

	err := r.db.Where("org_id = ?", id).
		Find(&tCamps).Error
	if err != nil {
		return nil, err
	}

	domain := []*domain.TCamp{}
	conv := converters.NewTCampConverter()
	for _, item := range tCamps {
		r, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, r)
	}
	return domain, nil
}

func (r *TCampRepository) RegisterSportsman(
	tCampApplication *domain.TCampApplication,
) error {
	model, err := converters.NewTCApplConverter().
		ToModel(tCampApplication)
	if err != nil {
		return err
	}

	err = r.db.Create(model).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TCampRepository) DeleteRegistration(smID, tCampID uuid.UUID) error {
	err := r.db.Where("sportsman_id = ? AND t_camp_id = ?", smID, tCampID).
		Delete(&models.TCampApplication{}).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dataaccess.ErrNotFound
		}
	}
	return nil
}

func (r *TCampRepository) GetUpcoming(smID uuid.UUID) (
	[]*domain.TCamp,
	error,
) {
	tCamps := []*models.TCamp{}
	if err := r.db.Raw("SELECT * FROM get_upcoming_tcamps_sm(?)", smID).
		Find(&tCamps).
		Error; err != nil {
		return nil, err
	}

	domain := []*domain.TCamp{}
	conv := converters.NewTCampConverter()
	for _, item := range tCamps {
		c, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, c)
	}

	return domain, nil
}

func (r *TCampRepository) ListUpcoming() ([]*domain.TCamp, error) {
	tCamps := []*models.TCamp{}
	if err := r.db.Raw("SELECT * FROM get_upcoming_tcamps()").
		Find(&tCamps).
		Error; err != nil {
		return nil, err
	}

	domain := []*domain.TCamp{}
	conv := converters.NewTCampConverter()
	for _, item := range tCamps {
		c, err := conv.ToDomain(item)
		if err != nil {
			return nil, err
		}
		domain = append(domain, c)
	}

	return domain, nil
}

func (r *TCampRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&models.TCamp{}, id).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
