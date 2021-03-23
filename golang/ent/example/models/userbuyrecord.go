package models

import (
	"context"

	"sqlent/ent"
)

type UserBuyRecordDao struct {
	conn *ent.Client
}

func NewUserBuyRecordDao(conn *ent.Client) *UserBuyRecordDao {
	return &UserBuyRecordDao{conn: conn}
}

func (d *UserBuyRecordDao) CreateUserBuyRecord(ctx context.Context, own *ent.User, ubr *ent.UserBuyRecord) (*ent.UserBuyRecord, error) {
	return d.conn.UserBuyRecord.
		Create().
		SetOwnerID(own.ID).
		SetPower(ubr.Power).
		SetPowerNum(ubr.PowerNum).
		SetTotalPower(ubr.TotalPower).
		SetTotalDay(ubr.TotalDay).
		SetRemainDay(ubr.RemainDay).
		SetNode(ubr.Node).
		SetUsedUsdt(ubr.UsedUsdt).
		SetBuyDate(ubr.BuyDate).
		Save(ctx)
}
