package consumers

import (
	"base_service/logger"
	"base_service/web/utils"
	"encoding/json"
	"fmt"
	"log/slog"

	pgjsonb "github.com/asif-mahmud/pg-jsonb"
	"github.com/guregu/null/v5"
)

const customerType = "Customer"

const userResourceTypeName = "User"
const userSource = "user-service:" + UserRoutingKey

type User struct {
	Id        int      `json:"Id" validate:"required"`
	CreatedBy null.Int `json:"CreatedBy" validate:"required_without=UpdatedBy"`
	UpdatedBy null.Int `json:"UpdatedBy" validate:"required_without=CreatedBy"`
	UserType  string   `json:"UserType"`
}

type UserLog struct {
	User     User
	JsonData pgjsonb.JSONB
}

func newUserLog(payload []byte) (*UserLog, error) {
	var userLog UserLog
	err := json.Unmarshal(payload, &userLog.User)
	if err != nil {
		return nil, err
	}
	err = userLog.JsonData.Scan(payload)
	if err != nil {
		return nil, err
	}
	return &userLog, nil
}

func (b *UserLog) ResourceTypeName() string {
	return userResourceTypeName
}

func (b *UserLog) ResourceID() string {
	return fmt.Sprintf("%d", b.User.Id)
}

func (b *UserLog) UserID() int {
	if b.User.UpdatedBy.Valid {
		return int(b.User.UpdatedBy.Int64)
	}

	if b.User.CreatedBy.Valid {
		return int(b.User.CreatedBy.Int64)
	}

	return 0
}

func (b *UserLog) Content() pgjsonb.JSONB {
	return b.JsonData
}

func (b *UserLog) Source() string {
	return userSource
}

func (b *UserLog) Validate() error {
	if err := utils.Validate(b.User); err != nil {
		return err
	}
	return nil
}

func ConsumeUser(data []byte) error {
	Log, err := newUserLog(data)
	if err != nil {
		slog.Error(
			"failed to create new user log",
			logger.Extra(map[string]any{
				"error":   err.Error(),
				"payload": string(data)},
			),
		)
		return err
	}
	fmt.Println(Log)

	if Log.User.UserType == customerType {
		return nil
	}

	// err = db.SaveLog(log)
	// if err != nil {
	// 	slog.Error(
	// 		"failed to save user log",
	// 		logger.Extra(map[string]any{
	// 			"error":   err.Error(),
	// 			"payload": string(data)},
	// 		),
	// 	)
	// 	return err
	// }
	return nil
}
