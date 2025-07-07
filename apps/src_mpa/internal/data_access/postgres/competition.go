package dataaccess

import (
	"errors"
	"src/internal/entities"
	"src/internal/service"

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

func (r *CompetitionRepository) Update(competition *entities.Competition) error {
	err := r.db.Save(competition).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return service.ErrNotFound
		}
		return err
	}
	return nil
}

func (r *CompetitionRepository) Create(competition *entities.Competition) error {
	competition.ID = uuid.New()
	err := r.db.Create(competition).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CompetitionRepository) GetCompetitionByID(competitionID uuid.UUID) (*entities.Competition,
	error) {
	competition := &entities.Competition{}
	err := r.db.Where("id = ?", competitionID).First(competition).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, err
	}
	return competition, nil
}

func (r *CompetitionRepository) ListCompetitions() ([]*entities.Competition, error) {
	competitions := []*entities.Competition{}
	err := r.db.Find(&competitions).Error
	if err != nil {
		return nil, err
	}
	return competitions, nil
}

func (r *CompetitionRepository) RegisterSportsman(compApplication *entities.CompApplication) error {
	err := r.db.Create(compApplication).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CompetitionRepository) DeleteRegistration(smID, compID uuid.UUID) error {
	err := r.db.Where("sportsman_id = ? AND competition_id = ?", smID, compID).Delete(&entities.CompApplication{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return service.ErrNotFound
		}
	}
	return nil
}

func (r *CompetitionRepository) GetUpcoming(smID uuid.UUID) ([]*entities.Competition, error) {
	comps := []*entities.Competition{}
	if err := r.db.Raw("SELECT * FROM get_upcoming_comps_sm(?)", smID).Find(&comps).Error; err != nil {
		return nil, err
	}

	return comps, nil
}
func (r *CompetitionRepository) ListUpcoming() ([]*entities.Competition, error) {
	comps := []*entities.Competition{}
	if err := r.db.Raw("SELECT * FROM get_upcoming_comps()").Find(&comps).Error; err != nil {
		return nil, err
	}

	return comps, nil
}
