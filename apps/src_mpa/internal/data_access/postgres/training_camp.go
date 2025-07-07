package dataaccess

import (
	"errors"
	"src/internal/entities"
	"src/internal/service"

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

func (r *TCampRepository) Update(tCamp *entities.TCamp) error {
	err := r.db.Save(tCamp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return service.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *TCampRepository) Create(tCamp *entities.TCamp) error {
	tCamp.ID = uuid.New()
	err := r.db.Create(tCamp).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TCampRepository) GetTCampByID(tCampID uuid.UUID) (*entities.TCamp, error) {
	tCamp := &entities.TCamp{}
	err := r.db.Where("id = ?", tCampID).First(tCamp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return tCamp, nil
}

func (r *TCampRepository) ListTCamps() ([]*entities.TCamp, error) {
	tCamps := []*entities.TCamp{}
	err := r.db.Find(&tCamps).Error
	if err != nil {
		return nil, err
	}
	return tCamps, nil
}

func (r *TCampRepository) RegisterSportsman(tCampApplication *entities.TCampApplication) error {
	err := r.db.Create(tCampApplication).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TCampRepository) DeleteRegistration(smID, tCampID uuid.UUID) error {
	err := r.db.Where("sportsman_id = ? AND t_camp_id = ?", smID, tCampID).Delete(&entities.TCampApplication{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return service.ErrNotFound
		}
	}
	return nil
}

func (r *TCampRepository) GetUpcoming(smID uuid.UUID) ([]*entities.TCamp, error) {
	tCamps := []*entities.TCamp{}
	if err := r.db.Raw("SELECT * FROM get_upcoming_tcamps_sm(?)", smID).Find(&tCamps).Error; err != nil {
		return nil, err
	}

	return tCamps, nil
}

func (r *TCampRepository) ListUpcoming() ([]*entities.TCamp, error) {
	tCamps := []*entities.TCamp{}
	if err := r.db.Raw("SELECT * FROM get_upcoming_tcamps()").Find(&tCamps).Error; err != nil {
		return nil, err
	}

	return tCamps, nil
}
