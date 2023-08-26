package infra

import "thresher/infra/model"

func InitDatabase() error {
	db := NewPostgresConnector()
	err := db.Conn.AutoMigrate(model.Users{}, model.Home{}, model.Encounter{}, model.Location{}, model.Posts{}, model.Message{})
	if err != nil {
		return err
	}
	return nil
}
