package data

import (
	"context"
	"dhb/app/app/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Address   string    `gorm:"type:varchar(100)"`
	Password  string    `gorm:"type:varchar(100)"`
	Undo      int64     `gorm:"type:int;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Trade struct {
	ID           int64     `gorm:"primarykey;type:int"`
	UserId       int64     `gorm:"type:int"`
	AmountCsd    int64     `gorm:"type:bigint"`
	RelAmountCsd int64     `gorm:"type:bigint"`
	AmountHbs    int64     `gorm:"type:bigint"`
	RelAmountHbs int64     `gorm:"type:bigint"`
	CsdReward    int64     `gorm:"type:bigint"`
	Status       string    `gorm:"type:varchar(45);not null"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
}

type UserInfo struct {
	ID               int64     `gorm:"primarykey;type:int"`
	UserId           int64     `gorm:"type:int;not null"`
	Vip              int64     `gorm:"type:int;not null"`
	UseVip           int64     `gorm:"type:int;not null"`
	HistoryRecommend int64     `gorm:"type:int;not null"`
	CreatedAt        time.Time `gorm:"type:datetime;not null"`
	UpdatedAt        time.Time `gorm:"type:datetime;not null"`
	TeamCsdBalance   int64     `gorm:"type:bigint;not null"`
}

type UserRecommend struct {
	ID            int64     `gorm:"primarykey;type:int"`
	UserId        int64     `gorm:"type:int;not null"`
	RecommendCode string    `gorm:"type:varchar(10000);not null"`
	CreatedAt     time.Time `gorm:"type:datetime;not null"`
	UpdatedAt     time.Time `gorm:"type:datetime;not null"`
}

type BalanceReward struct {
	ID             int64     `gorm:"primarykey;type:int"`
	UserId         int64     `gorm:"type:int;not null"`
	Amount         int64     `gorm:"type:bigint;not null"`
	Status         int64     `gorm:"type:int;not null"`
	H              int64     `gorm:"type:int;not null"`
	M              int64     `gorm:"type:int;not null"`
	SetDate        time.Time `gorm:"type:datetime;not null"`
	LastRewardDate time.Time `gorm:"type:datetime;not null"`
	CreatedAt      time.Time `gorm:"type:datetime;not null"`
	UpdatedAt      time.Time `gorm:"type:datetime;not null"`
}

type UserRecommendArea struct {
	ID            int64     `gorm:"primarykey;type:int"`
	RecommendCode string    `gorm:"type:varchar(10000);not null"`
	Version       int64     `gorm:"type:int;not null"`
	Num           int64     `gorm:"type:int;not null"`
	CreatedAt     time.Time `gorm:"type:datetime;not null"`
	UpdatedAt     time.Time `gorm:"type:datetime;not null"`
}

type UserArea struct {
	ID         int64     `gorm:"primarykey;type:int"`
	UserId     int64     `gorm:"type:int;not null"`
	Level      int64     `gorm:"type:int;not null"`
	Amount     int64     `gorm:"type:bigint;not null"`
	SelfAmount int64     `gorm:"type:bigint;not null"`
	CreatedAt  time.Time `gorm:"type:datetime;not null"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null"`
}

type UserCurrentMonthRecommend struct {
	ID              int64     `gorm:"primarykey;type:int"`
	UserId          int64     `gorm:"type:int;not null"`
	RecommendUserId int64     `gorm:"type:int;not null"`
	Date            time.Time `gorm:"type:datetime;not null"`
	CreatedAt       time.Time `gorm:"type:datetime;not null"`
	UpdatedAt       time.Time `gorm:"type:datetime;not null"`
}

type Config struct {
	ID        int64     `gorm:"primarykey;type:int"`
	Name      string    `gorm:"type:varchar(45);not null"`
	KeyName   string    `gorm:"type:varchar(45);not null"`
	Value     string    `gorm:"type:varchar(1000);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBalance struct {
	ID          int64     `gorm:"primarykey;type:int"`
	UserId      int64     `gorm:"type:int"`
	BalanceUsdt int64     `gorm:"type:bigint"`
	BalanceDhb  int64     `gorm:"type:bigint"`
	CreatedAt   time.Time `gorm:"type:datetime;not null"`
	UpdatedAt   time.Time `gorm:"type:datetime;not null"`
}

type Withdraw struct {
	ID              int64     `gorm:"primarykey;type:int"`
	UserId          int64     `gorm:"type:int"`
	Amount          int64     `gorm:"type:bigint"`
	RelAmount       int64     `gorm:"type:bigint"`
	Status          string    `gorm:"type:varchar(45);not null"`
	Type            string    `gorm:"type:varchar(45);not null"`
	BalanceRecordId int64     `gorm:"type:int"`
	CreatedAt       time.Time `gorm:"type:datetime;not null"`
	UpdatedAt       time.Time `gorm:"type:datetime;not null"`
}

type UserBalanceRecord struct {
	ID        int64     `gorm:"primarykey;type:int"`
	UserId    int64     `gorm:"type:int"`
	Balance   int64     `gorm:"type:bigint"`
	Amount    int64     `gorm:"type:bigint"`
	Type      string    `gorm:"type:varchar(45);not null"`
	CoinType  string    `gorm:"type:varchar(45);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Reward struct {
	ID               int64     `gorm:"primarykey;type:int"`
	UserId           int64     `gorm:"type:int;not null"`
	Amount           int64     `gorm:"type:bigint;not null"`
	AmountB          int64     `gorm:"type:bigint;not null"`
	BalanceRecordId  int64     `gorm:"type:int;not null"`
	Type             string    `gorm:"type:varchar(45);not null"`
	TypeRecordId     int64     `gorm:"type:int;not null"`
	Reason           string    `gorm:"type:varchar(45);not null"`
	ReasonLocationId int64     `gorm:"type:int;not null"`
	LocationType     string    `gorm:"type:varchar(45);not null"`
	CreatedAt        time.Time `gorm:"type:datetime;not null"`
	UpdatedAt        time.Time `gorm:"type:datetime;not null"`
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

type ConfigRepo struct {
	data *Data
	log  *log.Helper
}

type UserInfoRepo struct {
	data *Data
	log  *log.Helper
}

type UserRecommendRepo struct {
	data *Data
	log  *log.Helper
}

type UserCurrentMonthRecommendRepo struct {
	data *Data
	log  *log.Helper
}

type UserBalanceRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserInfoRepo(data *Data, logger log.Logger) biz.UserInfoRepo {
	return &UserInfoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewConfigRepo(data *Data, logger log.Logger) biz.ConfigRepo {
	return &ConfigRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserBalanceRepo(data *Data, logger log.Logger) biz.UserBalanceRepo {
	return &UserBalanceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserRecommendRepo(data *Data, logger log.Logger) biz.UserRecommendRepo {
	return &UserRecommendRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func NewUserCurrentMonthRecommendRepo(data *Data, logger log.Logger) biz.UserCurrentMonthRecommendRepo {
	return &UserCurrentMonthRecommendRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// GetUserByAddress .
func (u *UserRepo) GetUserByAddress(ctx context.Context, address string) (*biz.User, error) {
	var user User
	if err := u.data.db.Where(&User{Address: address}).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	return &biz.User{
		ID:       user.ID,
		Address:  user.Address,
		Password: user.Password,
	}, nil
}

// GetConfigByKeys .
func (c *ConfigRepo) GetConfigByKeys(ctx context.Context, keys ...string) ([]*biz.Config, error) {
	var configs []*Config
	res := make([]*biz.Config, 0)
	if err := c.data.db.Where("key_name IN (?)", keys).Table("config").Find(&configs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("CONFIG_NOT_FOUND", "config not found")
		}

		return nil, errors.New(500, "Config ERROR", err.Error())
	}

	for _, config := range configs {
		res = append(res, &biz.Config{
			ID:      config.ID,
			KeyName: config.KeyName,
			Name:    config.Name,
			Value:   config.Value,
		})
	}

	return res, nil
}

// GetConfigs .
func (c *ConfigRepo) GetConfigs(ctx context.Context) ([]*biz.Config, error) {
	var configs []*Config
	res := make([]*biz.Config, 0)
	if err := c.data.db.Table("config").Find(&configs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("CONFIG_NOT_FOUND", "config not found")
		}

		return nil, errors.New(500, "Config ERROR", err.Error())
	}

	for _, config := range configs {
		res = append(res, &biz.Config{
			ID:      config.ID,
			KeyName: config.KeyName,
			Name:    config.Name,
			Value:   config.Value,
		})
	}

	return res, nil
}

// UpdateConfig .
func (c *ConfigRepo) UpdateConfig(ctx context.Context, id int64, value string) (bool, error) {
	var config Config
	config.Value = value

	res := c.data.DB(ctx).Table("config").Where("id=?", id).Updates(&config)
	if res.Error != nil {
		return false, errors.New(500, "UPDATE_USER_INFO_ERROR", "用户信息修改失败")
	}

	return true, nil
}

// GetUserById .
func (u *UserRepo) GetUserById(ctx context.Context, Id int64) (*biz.User, error) {
	var user User
	if err := u.data.db.Where(&User{ID: Id}).Table("user").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	return &biz.User{
		ID:       user.ID,
		Password: user.Password,
		Address:  user.Address,
		Undo:     user.Undo,
	}, nil
}

// GetUserInfoByUserId .
func (ui *UserInfoRepo) GetUserInfoByUserId(ctx context.Context, userId int64) (*biz.UserInfo, error) {
	var userInfo UserInfo
	if err := ui.data.db.Where(&UserInfo{UserId: userId}).Table("user_info").First(&userInfo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USERINFO_NOT_FOUND", "userinfo not found")
		}

		return nil, errors.New(500, "USERINFO ERROR", err.Error())
	}

	return &biz.UserInfo{
		ID:               userInfo.ID,
		UserId:           userInfo.UserId,
		Vip:              userInfo.Vip,
		HistoryRecommend: userInfo.HistoryRecommend,
		TeamCsdBalance:   userInfo.TeamCsdBalance,
		UseVip:           userInfo.UseVip,
	}, nil
}

// GetUserByAddresses .
func (u *UserRepo) GetUserByAddresses(ctx context.Context, Addresses ...string) (map[string]*biz.User, error) {
	var users []*User
	if err := u.data.db.Table("user").Where("address IN (?)", Addresses).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	res := make(map[string]*biz.User, 0)
	for _, item := range users {
		res[item.Address] = &biz.User{
			ID:      item.ID,
			Address: item.Address,
		}
	}
	return res, nil
}

// GetUserCount .
func (u *UserRepo) GetUserCount(ctx context.Context) (int64, error) {
	var count int64
	if err := u.data.db.Table("user").Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return count, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return count, errors.New(500, "USER ERROR", err.Error())
	}

	return count, nil
}

// GetUserCountToday .
func (u *UserRepo) GetUserCountToday(ctx context.Context) (int64, error) {
	var count int64
	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 14 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 15, 59, 59, 0, time.UTC)

	if err := u.data.db.Table("user").
		Where("created_at>=?", todayStart).Where("created_at<=?", todayEnd).Count(&count).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return count, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return count, errors.New(500, "USER ERROR", err.Error())
	}

	return count, nil
}

// GetUserByUserIds .
func (u *UserRepo) GetUserByUserIds(ctx context.Context, userIds ...int64) (map[int64]*biz.User, error) {
	var users []*User
	if err := u.data.db.Table("user").Where("id IN (?)", userIds).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found")
		}

		return nil, errors.New(500, "USER ERROR", err.Error())
	}

	res := make(map[int64]*biz.User, 0)
	for _, item := range users {
		res[item.ID] = &biz.User{
			ID:      item.ID,
			Address: item.Address,
		}
	}
	return res, nil
}

// GetUsers .
func (u *UserRepo) GetUsers(ctx context.Context, b *biz.Pagination, address string) ([]*biz.User, error, int64) {
	var (
		users []*User
		count int64
	)
	instance := u.data.db.Table("user")

	if "" != address {
		instance = instance.Where("address=?", address)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Order("id desc").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_NOT_FOUND", "user not found"), 0
		}

		return nil, errors.New(500, "USER ERROR", err.Error()), 0
	}

	res := make([]*biz.User, 0)
	for _, item := range users {
		res = append(res, &biz.User{
			ID:        item.ID,
			Address:   item.Address,
			CreatedAt: item.CreatedAt,
		})
	}
	return res, nil, count
}

// CreateUser .
func (u *UserRepo) CreateUser(ctx context.Context, uc *biz.User) (*biz.User, error) {
	var user User
	user.Address = uc.Address
	user.Password = uc.Password
	if 0 == len(uc.Password) {
		user.Password = "zkf123"
	}
	res := u.data.DB(ctx).Table("user").Create(&user)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ERROR", "用户创建失败")
	}

	return &biz.User{
		ID:       user.ID,
		Address:  user.Address,
		Password: user.Password,
	}, nil
}

// CreateUserInfo .
func (ui *UserInfoRepo) CreateUserInfo(ctx context.Context, u *biz.User) (*biz.UserInfo, error) {
	var userInfo UserInfo
	userInfo.UserId = u.ID

	res := ui.data.DB(ctx).Table("user_info").Create(&userInfo)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_INFO_ERROR", "用户信息创建失败")
	}

	return &biz.UserInfo{
		ID:               userInfo.ID,
		UserId:           userInfo.UserId,
		Vip:              0,
		HistoryRecommend: 0,
	}, nil
}

// UpdateUserInfo .
func (ui *UserInfoRepo) UpdateUserInfo(ctx context.Context, u *biz.UserInfo) (*biz.UserInfo, error) {
	var userInfo UserInfo
	userInfo.Vip = u.Vip
	userInfo.HistoryRecommend = u.HistoryRecommend

	res := ui.data.DB(ctx).Table("user_info").Where("user_id=?", u.UserId).Updates(&userInfo)
	if res.Error != nil {
		return nil, errors.New(500, "UPDATE_USER_INFO_ERROR", "用户信息修改失败")
	}

	return &biz.UserInfo{
		ID:               userInfo.ID,
		UserId:           userInfo.UserId,
		Vip:              userInfo.Vip,
		HistoryRecommend: userInfo.HistoryRecommend,
	}, nil
}

// GetUserRecommendByUserId .
func (ur *UserRecommendRepo) GetUserRecommendByUserId(ctx context.Context, userId int64) (*biz.UserRecommend, error) {
	var userRecommend UserRecommend
	if err := ur.data.db.Where(&UserRecommend{UserId: userId}).Table("user_recommend").First(&userRecommend).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	return &biz.UserRecommend{
		UserId:        userRecommend.UserId,
		RecommendCode: userRecommend.RecommendCode,
	}, nil
}

// GetUserRecommendByCode .
func (ur *UserRecommendRepo) GetUserRecommendByCode(ctx context.Context, code string) ([]*biz.UserRecommend, error) {
	var (
		userRecommends []*UserRecommend
	)
	res := make([]*biz.UserRecommend, 0)

	instance := ur.data.db.Table("user_recommend").Where("recommend_code=?", code)

	if err := instance.Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
			CreatedAt:     userRecommend.CreatedAt,
		})
	}

	return res, nil
}

// GetUserRecommendLikeCode .
func (ur *UserRecommendRepo) GetUserRecommendLikeCode(ctx context.Context, code string) ([]*biz.UserRecommend, error) {
	var userRecommends []*UserRecommend
	res := make([]*biz.UserRecommend, 0)
	if err := ur.data.db.Where("recommend_code Like ?", code+"%").Table("user_recommend").Find(&userRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_RECOMMEND_NOT_FOUND", "user recommend not found")
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	for _, userRecommend := range userRecommends {
		res = append(res, &biz.UserRecommend{
			UserId:        userRecommend.UserId,
			RecommendCode: userRecommend.RecommendCode,
		})
	}

	return res, nil
}

// CreateUserRecommend .
func (ur *UserRecommendRepo) CreateUserRecommend(ctx context.Context, u *biz.User, recommendUser *biz.UserRecommend) (*biz.UserRecommend, error) {
	var tmpRecommendCode string
	if nil != recommendUser && 0 < recommendUser.UserId {
		tmpRecommendCode = "D" + strconv.FormatInt(recommendUser.UserId, 10)
		if "" != recommendUser.RecommendCode {
			tmpRecommendCode = recommendUser.RecommendCode + tmpRecommendCode
		}
	}

	var userRecommend UserRecommend
	userRecommend.UserId = u.ID
	userRecommend.RecommendCode = tmpRecommendCode

	res := ur.data.DB(ctx).Table("user_recommend").Create(&userRecommend)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_RECOMMEND_ERROR", "用户推荐关系创建失败")
	}

	return &biz.UserRecommend{
		ID:            userRecommend.ID,
		UserId:        userRecommend.UserId,
		RecommendCode: userRecommend.RecommendCode,
	}, nil
}

// CreateUserRecommendArea .
func (ur *UserRecommendRepo) CreateUserRecommendArea(ctx context.Context, u *biz.User, recommendUser *biz.UserRecommend) (bool, error) {
	var tmpRecommendCode string
	if nil != recommendUser && 0 < recommendUser.UserId {
		tmpRecommendCode = "D" + strconv.FormatInt(recommendUser.UserId, 10)
		if "" != recommendUser.RecommendCode {
			tmpRecommendCode = recommendUser.RecommendCode + tmpRecommendCode
		}
	}

	var myUserRecommendArea UserRecommendArea
	if err := ur.data.db.Where("recommend_code=?", tmpRecommendCode).Table("user_recommend_area").First(&myUserRecommendArea).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, errors.New(500, "USER RECOMMEND ERROR", err.Error())
		}

		// 业务上限制了错误的上一级未insert下一级优先insert的情况
		var userRecommendArea UserRecommendArea
		userRecommendArea.RecommendCode = tmpRecommendCode + "D" + strconv.FormatInt(u.ID, 10)
		userRecommendArea.Num = myUserRecommendArea.Num + int64(len(strings.Split(userRecommendArea.RecommendCode, "D"))-1)
		res := ur.data.DB(ctx).Table("user_recommend_area").Create(&userRecommendArea)
		if res.Error != nil {
			return false, errors.New(500, "CREATE_USER_RECOMMEND_AREA_ERROR", "用户推荐关系链路创建失败")
		}

	} else {
		res := ur.data.DB(ctx).Table("user_recommend_area").
			Where("id=? and version=?", myUserRecommendArea.ID, myUserRecommendArea.Version).
			Updates(map[string]interface{}{"version": gorm.Expr("version + ?", 1), "num": gorm.Expr("num + ?", 1), "recommend_code": tmpRecommendCode + "D" + strconv.FormatInt(u.ID, 10)})
		if 0 == res.RowsAffected || nil != res.Error {
			return false, errors.New(500, "CREATE_USER_RECOMMEND_AREA_ERROR", "用户推荐关系链路修改失败")
		}
	}

	return true, nil
}

// CreateUserArea .
func (ur *UserRecommendRepo) CreateUserArea(ctx context.Context, u *biz.User) (bool, error) {
	// 业务上限制了错误的上一级未insert下一级优先insert的情况
	var userArea UserArea
	userArea.UserId = u.ID
	res := ur.data.DB(ctx).Table("user_area").Create(&userArea)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_AREA_ERROR", "用户区信息创建失败")
	}

	return true, nil
}

// DeleteOrOriginUserRecommendArea .
func (ur *UserRecommendRepo) DeleteOrOriginUserRecommendArea(ctx context.Context, code string, originCode string) (bool, error) {
	//var myUserRecommendArea []*UserRecommendArea
	//if err := ur.data.db.Where("recommend_code like ? and status=?", originCode+"%", 0).Table("user_recommend_area").Find(&myUserRecommendArea).Error; err != nil {
	//	if !errors.Is(err, gorm.ErrRecordNotFound) {
	//		return false, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	//	}
	//
	//	res := ur.data.DB(ctx).Table("user_recommend_area").
	//		Where("recommend_code=? and status=?", code, 0).
	//		Updates(map[string]interface{}{"recommend_code": originCode})
	//	if 0 == res.RowsAffected || nil != res.Error {
	//		return false, errors.New(500, "UPDATE_USER_RECOMMEND_AREA_ERROR", "用户推荐关系链路修改失败")
	//	}
	//	return true, nil
	//}
	//
	//if 2 > len(myUserRecommendArea) {
	res := ur.data.DB(ctx).Table("user_recommend_area").
		Where("recommend_code=?", code).
		Updates(map[string]interface{}{"recommend_code": originCode, "num": gorm.Expr("num - ?", 1)})
	if 0 == res.RowsAffected || nil != res.Error {
		return false, errors.New(500, "UPDATE_USER_RECOMMEND_AREA_ERROR", "用户推荐关系链路修改失败")
	}
	//} else {
	//	res := ur.data.DB(ctx).Table("user_recommend_area").
	//		Where("recommend_code=? and status=?", code, 0).
	//		Updates(map[string]interface{}{"status": 1})
	//	if 0 == res.RowsAffected || nil != res.Error {
	//		return false, errors.New(500, "UPDATE_USER_RECOMMEND_AREA_ERROR", "用户推荐关系链路修改失败")
	//	}
	//}

	return true, nil
}

// GetUserRecommendLowArea .
func (ur *UserRecommendRepo) GetUserRecommendLowArea(ctx context.Context, code string) ([]*biz.UserRecommendArea, error) {

	var firstRecommendArea *UserRecommendArea
	if err := ur.data.db.Order("num desc").Table("user_recommend_area").First(&firstRecommendArea).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(500, "USER RECOMMEND NOT FOUND", err.Error())
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	var myUserRecommendAreas []*UserRecommendArea
	if err := ur.data.db.Where("recommend_code like ?", code+"%").Table("user_recommend_area").Find(&myUserRecommendAreas).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(500, "USER RECOMMEND NOT FOUND", err.Error())
		}

		return nil, errors.New(500, "USER RECOMMEND ERROR", err.Error())
	}

	res := make([]*biz.UserRecommendArea, 0)
	for _, v := range myUserRecommendAreas {
		if firstRecommendArea.ID == v.ID {
			continue
		}

		res = append(res, &biz.UserRecommendArea{
			ID:            v.ID,
			RecommendCode: v.RecommendCode,
			Num:           v.Num,
		})
	}

	return res, nil
}

// GetUserArea .
func (ur *UserRecommendRepo) GetUserArea(ctx context.Context, userId int64) (*biz.UserArea, error) {

	var userArea *UserArea
	if err := ur.data.db.Where("user_id=?", userId).Table("user_area").First(&userArea).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(500, "USER AREA NOT FOUND", err.Error())
		}

		return nil, errors.New(500, "USER AREA ERROR", err.Error())
	}

	return &biz.UserArea{
		ID:         userArea.ID,
		UserId:     userArea.UserId,
		Amount:     userArea.Amount,
		SelfAmount: userArea.SelfAmount,
		Level:      userArea.Level,
	}, nil
}

// GetUserAreas .
func (ur *UserRecommendRepo) GetUserAreas(ctx context.Context, userIds []int64) ([]*biz.UserArea, error) {

	var userAreas []*UserArea
	if err := ur.data.db.Where("user_id in (?)", userIds).Table("user_area").Find(&userAreas).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(500, "USER AREA NOT FOUND", err.Error())
		}

		return nil, errors.New(500, "USER AREA ERROR", err.Error())
	}

	res := make([]*biz.UserArea, 0)
	for _, v := range userAreas {
		res = append(res, &biz.UserArea{
			ID:         v.ID,
			UserId:     v.UserId,
			Amount:     v.Amount,
			SelfAmount: v.SelfAmount,
			Level:      v.Level,
		})
	}

	return res, nil
}

// UpdateUserRecommend .
func (ur *UserRecommendRepo) UpdateUserRecommend(ctx context.Context, u *biz.User, recommendUser *biz.UserRecommend) (bool, error) {
	var tmpRecommendCode string
	if nil != recommendUser && 0 < recommendUser.UserId {
		tmpRecommendCode = "D" + strconv.FormatInt(recommendUser.UserId, 10)
		if "" != recommendUser.RecommendCode {
			tmpRecommendCode = recommendUser.RecommendCode + tmpRecommendCode
		}
	}

	var userRecommend UserRecommend
	userRecommend.RecommendCode = tmpRecommendCode

	res := ur.data.DB(ctx).Table("user_recommend").Where("user_id", u.ID).Updates(&userRecommend)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_RECOMMEND_ERROR", "用户推荐关系修改失败")
	}

	return true, nil
}

// CreateUserBalance .
func (ub UserBalanceRepo) CreateUserBalance(ctx context.Context, u *biz.User) (*biz.UserBalance, error) {
	var userBalance UserBalance
	userBalance.UserId = u.ID
	res := ub.data.DB(ctx).Table("user_balance").Create(&userBalance)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_BALANCE_ERROR", "用户余额信息创建失败")
	}

	return &biz.UserBalance{
		ID:          userBalance.ID,
		UserId:      userBalance.UserId,
		BalanceUsdt: userBalance.BalanceUsdt,
		BalanceDhb:  userBalance.BalanceDhb,
	}, nil
}

// CreateUserBalanceLock .
func (ub UserBalanceRepo) CreateUserBalanceLock(ctx context.Context, u *biz.User) (*biz.UserBalance, error) {
	var userBalance UserBalance
	userBalance.UserId = u.ID
	res := ub.data.DB(ctx).Table("user_balance_lock").Create(&userBalance)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_BALANCE_ERROR", "用户余额信息创建失败")
	}

	return &biz.UserBalance{
		ID:          userBalance.ID,
		UserId:      userBalance.UserId,
		BalanceUsdt: userBalance.BalanceUsdt,
		BalanceDhb:  userBalance.BalanceDhb,
	}, nil
}

// GetUserBalance .
func (ub UserBalanceRepo) GetUserBalance(ctx context.Context, userId int64) (*biz.UserBalance, error) {
	var userBalance UserBalance
	if err := ub.data.db.Where("user_id=?", userId).Table("user_balance").First(&userBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return nil, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return &biz.UserBalance{
		ID:          userBalance.ID,
		UserId:      userBalance.UserId,
		BalanceUsdt: userBalance.BalanceUsdt,
		BalanceDhb:  userBalance.BalanceDhb,
	}, nil
}

// GetUserBalanceLock .
func (ub UserBalanceRepo) GetUserBalanceLock(ctx context.Context, userId int64) (*biz.UserBalance, error) {
	var userBalance UserBalance
	if err := ub.data.db.Where("user_id=?", userId).Table("user_balance_lock").First(&userBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return nil, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return &biz.UserBalance{
		ID:          userBalance.ID,
		UserId:      userBalance.UserId,
		BalanceUsdt: userBalance.BalanceUsdt,
		BalanceDhb:  userBalance.BalanceDhb,
	}, nil
}

// Trade .
func (ub *UserBalanceRepo) Trade(ctx context.Context, userId int64, amount int64, amountB int64, amountRel int64, amountBRel int64, tmpRecommendUserIdsInt []int64, amount2 int64) error {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance_lock").
		Where("user_id=? and balance_usdt>=?", userId, amount).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt - ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance_dhb>=?", userId, amountB-amountBRel).
		Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb - ?", amountB-amountBRel)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "trade"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	var userBalanceRecode1 UserBalanceRecord
	userBalanceRecode1.Balance = userBalance.BalanceDhb
	userBalanceRecode1.UserId = userBalance.UserId
	userBalanceRecode1.Type = "trade_dhb"
	userBalanceRecode1.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode1).Error
	if err != nil {
		return err
	}

	var trade Trade
	trade.AmountCsd = amount
	trade.AmountHbs = amountB
	trade.RelAmountCsd = amountRel
	trade.RelAmountHbs = amountBRel
	trade.UserId = userId
	trade.Status = "default"
	err = ub.data.DB(ctx).Table("trade").Create(&trade).Error
	if err != nil {
		return err
	}

	if 0 < len(tmpRecommendUserIdsInt) {
		if err = ub.data.DB(ctx).Table("user_info").
			Where("user_id in (?)", tmpRecommendUserIdsInt).
			Updates(map[string]interface{}{"team_csd_balance": gorm.Expr("team_csd_balance - ?", amount-amountRel)}).Error; nil != err {
			return errors.NotFound("user balance err", "user balance not found")
		}
	}

	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amountRel)}).Error; nil != err {
		return errors.NotFound("user balance err", "user balance not found")
	}

	return nil
}

// LocationReward .
func (ub *UserBalanceRepo) LocationReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "location" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "location" // 给我分红的理由
	reward.ReasonLocationId = myLocationId
	reward.LocationType = locationType
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// WithdrawReward .
func (ub *UserBalanceRepo) WithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64, myLocationId int64, locationType string) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "location" // 给我分红的理由
	reward.ReasonLocationId = myLocationId
	reward.LocationType = locationType
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// Deposit .
func (ub *UserBalanceRepo) Deposit(ctx context.Context, userId int64, amount int64) (int64, error) {
	var err error
	//if err = ub.data.DB(ctx).Table("user_balance").
	//	Where("user_id=?", userId).
	//	Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
	//	return 0, errors.NotFound("user balance err", "user balance not found")
	//}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "deposit"
	userBalanceRecode.CoinType = "usdt"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// DepositLast .
func (ub *UserBalanceRepo) DepositLast(ctx context.Context, userId int64, lastAmount int64, locationId int64) (int64, error) {
	var (
		err error
	)
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", lastAmount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = lastAmount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = lastAmount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "last"          // 本次分红的行为类型
	reward.Reason = "last_reward" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	res := ub.data.db.Table("location").
		Where("id=?", locationId).
		Updates(map[string]interface{}{"stop_location_again": "1"})
	if 0 == res.RowsAffected || res.Error != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// DepositDhb .
func (ub *UserBalanceRepo) DepositDhb(ctx context.Context, userId int64, amount int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceDhb
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "deposit"
	userBalanceRecode.CoinType = "dhb"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// UpdateWithdrawAmount .
func (ub *UserBalanceRepo) UpdateWithdrawAmount(ctx context.Context, id int64, status string, amount int64) (*biz.Withdraw, error) {
	var withdraw Withdraw
	withdraw.Status = status
	withdraw.Amount = amount
	res := ub.data.DB(ctx).Table("withdraw").Where("id=?", id).Updates(&withdraw)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_WITHDRAW_ERROR", "提现记录修改失败")
	}

	return &biz.Withdraw{
		ID:              withdraw.ID,
		UserId:          withdraw.UserId,
		Amount:          withdraw.Amount,
		RelAmount:       withdraw.RelAmount,
		BalanceRecordId: withdraw.BalanceRecordId,
		Status:          withdraw.Status,
		Type:            withdraw.Type,
		CreatedAt:       withdraw.CreatedAt,
	}, nil
}

// WithdrawUsdt .
func (ub *UserBalanceRepo) WithdrawUsdt(ctx context.Context, userId int64, amount int64, tmpRecommendUserIdsInt []int64) error {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance_usdt>=?", userId, amount).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt - ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	if 0 < len(tmpRecommendUserIdsInt) {
		if err = ub.data.DB(ctx).Table("user_info").
			Where("user_id in (?)", tmpRecommendUserIdsInt).
			Updates(map[string]interface{}{"team_csd_balance": gorm.Expr("team_csd_balance - ?", amount)}).Error; nil != err {
			return errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "withdraw"
	userBalanceRecode.CoinType = "usdt"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	return nil
}

// TranUsdt .
func (ub *UserBalanceRepo) TranUsdt(ctx context.Context, userId int64, toUserId int64, amount int64, tmpRecommendUserIdsInt []int64, tmpRecommendUserIdsInt2 []int64) error {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance_usdt>=?", userId, amount).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt - ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", toUserId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
		return errors.NotFound("user balance err", "user balance not found")
	}

	if len(tmpRecommendUserIdsInt2) > 0 {
		if err = ub.data.DB(ctx).Table("user_info").
			Where("user_id in (?)", tmpRecommendUserIdsInt2).
			Updates(map[string]interface{}{"team_csd_balance": gorm.Expr("team_csd_balance + ?", amount)}).Error; nil != err {
			return errors.NotFound("user balance err", "user balance not found")
		}
	}
	if 0 < len(tmpRecommendUserIdsInt) {
		if err = ub.data.DB(ctx).Table("user_info").
			Where("user_id in (?)", tmpRecommendUserIdsInt).
			Updates(map[string]interface{}{"team_csd_balance": gorm.Expr("team_csd_balance - ?", amount)}).Error; nil != err {
			return errors.NotFound("user balance err", "user balance not found")
		}
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "tran"
	userBalanceRecode.CoinType = "usdt"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	var userBalanceRecode1 UserBalanceRecord
	userBalanceRecode1.UserId = toUserId
	userBalanceRecode1.Type = "tran_to"
	userBalanceRecode1.CoinType = "usdt"
	userBalanceRecode1.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode1).Error
	if err != nil {
		return err
	}

	return nil
}

func (ub *UserBalanceRepo) TradeUsdt(ctx context.Context, userId int64, amount int64) error {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance_usdt>=?", userId, amount).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt - ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "withdraw"
	userBalanceRecode.CoinType = "usdt"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	return nil
}

// SetBalanceReward .
func (ub *UserBalanceRepo) SetBalanceReward(ctx context.Context, userId int64, amount int64) error {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance_usdt>=?", userId, amount).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt - ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	now := time.Now().UTC()
	var balanceReward BalanceReward
	balanceReward.Amount = amount
	balanceReward.UserId = userId
	balanceReward.SetDate = now
	balanceReward.Status = 1
	balanceReward.LastRewardDate = now
	balanceReward.H = int64(now.Hour())
	balanceReward.M = int64(now.Minute())
	err = ub.data.DB(ctx).Table("balance_reward").Create(&balanceReward).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateBalanceReward .
func (ub *UserBalanceRepo) UpdateBalanceReward(ctx context.Context, userId int64, id int64, amount int64, status int64) error {
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	if res := ub.data.DB(ctx).Table("balance_reward").
		Where("id=? and status=?", id, 1).
		Updates(map[string]interface{}{"amount": gorm.Expr("amount - ?", amount), "status": status}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	return nil
}

// WithdrawDhb .
func (ub *UserBalanceRepo) WithdrawDhb(ctx context.Context, userId int64, amount int64) error {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance_dhb>=?", userId, amount).
		Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb - ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceDhb
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "withdraw"
	userBalanceRecode.CoinType = "dhb"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	return nil
}

// TranDhb .
func (ub *UserBalanceRepo) TranDhb(ctx context.Context, userId int64, toUserId int64, amount int64) error {
	var err error
	if res := ub.data.DB(ctx).Table("user_balance").
		Where("user_id=? and balance_dhb>=?", userId, amount).
		Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb - ?", amount)}); 0 == res.RowsAffected || nil != res.Error {
		return errors.NotFound("user balance err", "user balance error")
	}

	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", toUserId).
		Updates(map[string]interface{}{"balance_dhb": gorm.Expr("balance_dhb + ?", amount)}).Error; nil != err {
		return errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceDhb
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "tran"
	userBalanceRecode.CoinType = "dhb"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return err
	}

	var userBalanceRecode1 UserBalanceRecord
	userBalanceRecode1.UserId = toUserId
	userBalanceRecode1.Type = "tran_to"
	userBalanceRecode1.CoinType = "dhb"
	userBalanceRecode1.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode1).Error
	if err != nil {
		return err
	}

	return nil
}

// GreateWithdraw .
func (ub *UserBalanceRepo) GreateWithdraw(ctx context.Context, userId int64, amount int64, amountFee int64, coinType string) (*biz.Withdraw, error) {
	var withdraw Withdraw
	withdraw.UserId = userId
	withdraw.Amount = amount
	withdraw.Type = coinType
	withdraw.Status = "rewarded"
	res := ub.data.DB(ctx).Table("withdraw").Create(&withdraw)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_WITHDRAW_ERROR", "提现记录创建失败")
	}

	var (
		reward Reward
		err    error
	)

	reward.UserId = 999999999
	reward.Amount = amountFee
	reward.BalanceRecordId = 999999999
	reward.Type = "withdraw" // 本次分红的行为类型
	if "dhb" == coinType {
		reward.Type = "withdraw_dhb" // 本次分红的行为类型
	}
	reward.TypeRecordId = withdraw.ID
	reward.Reason = "system_reward" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return nil, errors.New(500, "CREATE_WITHDRAW_ERROR", "提现记录创建失败")
	}

	return &biz.Withdraw{
		ID:              withdraw.ID,
		UserId:          withdraw.UserId,
		Amount:          withdraw.Amount,
		RelAmount:       withdraw.RelAmount,
		BalanceRecordId: withdraw.BalanceRecordId,
		Status:          withdraw.Status,
		Type:            withdraw.Type,
		CreatedAt:       withdraw.CreatedAt,
	}, nil
}

// UpdateWithdraw .
func (ub *UserBalanceRepo) UpdateWithdraw(ctx context.Context, id int64, status string) (*biz.Withdraw, error) {
	var withdraw Withdraw
	withdraw.Status = status
	res := ub.data.DB(ctx).Table("withdraw").Where("id=?", id).Updates(&withdraw)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_WITHDRAW_ERROR", "提现记录修改失败")
	}

	return &biz.Withdraw{
		ID:              withdraw.ID,
		UserId:          withdraw.UserId,
		Amount:          withdraw.Amount,
		RelAmount:       withdraw.RelAmount,
		BalanceRecordId: withdraw.BalanceRecordId,
		Status:          withdraw.Status,
		Type:            withdraw.Type,
		CreatedAt:       withdraw.CreatedAt,
	}, nil
}

// GetWithdrawByUserId .
func (ub *UserBalanceRepo) GetWithdrawByUserId(ctx context.Context, userId int64, typeCoin string) ([]*biz.Withdraw, error) {
	var withdraws []*Withdraw
	res := make([]*biz.Withdraw, 0)
	if err := ub.data.db.Where("user_id=?", userId).Where("type=?", typeCoin).Table("withdraw").Find(&withdraws).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, withdraw := range withdraws {
		res = append(res, &biz.Withdraw{
			ID:              withdraw.ID,
			UserId:          withdraw.UserId,
			Amount:          withdraw.Amount,
			RelAmount:       withdraw.RelAmount,
			BalanceRecordId: withdraw.BalanceRecordId,
			Status:          withdraw.Status,
			Type:            withdraw.Type,
			CreatedAt:       withdraw.CreatedAt,
		})
	}
	return res, nil
}

// GetUserBalanceRecordByUserId .
func (ub *UserBalanceRepo) GetUserBalanceRecordByUserId(ctx context.Context, userId int64, typeCoin string, tran string) ([]*biz.UserBalanceRecord, error) {
	var userBalanceRecord []*UserBalanceRecord
	res := make([]*biz.UserBalanceRecord, 0)
	if err := ub.data.db.Where("user_id=?", userId).
		Where("type=? and coin_type=?", tran, typeCoin).Table("user_balance_record").Find(&userBalanceRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, v := range userBalanceRecord {
		res = append(res, &biz.UserBalanceRecord{
			ID:        v.ID,
			UserId:    v.UserId,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

// GetUserBalanceRecordsByUserId .
func (ub *UserBalanceRepo) GetUserBalanceRecordsByUserId(ctx context.Context, userId int64) ([]*biz.UserBalanceRecord, error) {
	var userBalanceRecord []*UserBalanceRecord
	res := make([]*biz.UserBalanceRecord, 0)
	if err := ub.data.db.Where("user_id=?", userId).
		Where("type=? and (coin_type=? or coin_type=?)", "deposit", "CSD", "HBS").Table("user_balance_record").Find(&userBalanceRecord).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, v := range userBalanceRecord {
		res = append(res, &biz.UserBalanceRecord{
			ID:        v.ID,
			UserId:    v.UserId,
			Amount:    v.Amount,
			CoinType:  v.CoinType,
			CreatedAt: v.CreatedAt,
		})
	}
	return res, nil
}

// GetTradeByUserId .
func (ub *UserBalanceRepo) GetTradeByUserId(ctx context.Context, userId int64) ([]*biz.Trade, error) {
	var trades []*Trade
	res := make([]*biz.Trade, 0)
	if err := ub.data.db.Where("user_id=?", userId).Table("trade").Find(&trades).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, trade := range trades {
		res = append(res, &biz.Trade{
			ID:           trade.ID,
			UserId:       trade.UserId,
			AmountCsd:    trade.AmountCsd,
			RelAmountCsd: trade.RelAmountCsd,
			AmountHbs:    trade.AmountHbs,
			RelAmountHbs: trade.RelAmountHbs,
			Status:       trade.Status,
			CreatedAt:    trade.CreatedAt,
		})
	}
	return res, nil
}

// GetBalanceRewardByUserId .
func (ub *UserBalanceRepo) GetBalanceRewardByUserId(ctx context.Context, userId int64) ([]*biz.BalanceReward, error) {
	var balanceRewards []*BalanceReward
	res := make([]*biz.BalanceReward, 0)
	if err := ub.data.db.Where("user_id=?", userId).Where("status=?", 1).Order("id asc").Table("balance_reward").Find(&balanceRewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, balanceReward := range balanceRewards {
		res = append(res, &biz.BalanceReward{
			ID:      balanceReward.ID,
			UserId:  balanceReward.UserId,
			Status:  balanceReward.Status,
			Amount:  balanceReward.Amount,
			SetDate: balanceReward.SetDate,
		})
	}
	return res, nil
}

// GetWithdrawNotDeal .
func (ub *UserBalanceRepo) GetWithdrawNotDeal(ctx context.Context) ([]*biz.Withdraw, error) {
	var withdraws []*Withdraw
	res := make([]*biz.Withdraw, 0)
	if err := ub.data.db.Where("status=?", "").Table("withdraw").Find(&withdraws).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, withdraw := range withdraws {
		res = append(res, &biz.Withdraw{
			ID:              withdraw.ID,
			UserId:          withdraw.UserId,
			Amount:          withdraw.Amount,
			RelAmount:       withdraw.RelAmount,
			BalanceRecordId: withdraw.BalanceRecordId,
			Status:          withdraw.Status,
			Type:            withdraw.Type,
			CreatedAt:       withdraw.CreatedAt,
		})
	}
	return res, nil
}

func (ub *UserBalanceRepo) GetWithdrawById(ctx context.Context, id int64) (*biz.Withdraw, error) {
	var withdraw *Withdraw
	if err := ub.data.db.Where("id=?", id).Table("withdraw").First(&withdraw).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}
	return &biz.Withdraw{
		ID:              withdraw.ID,
		UserId:          withdraw.UserId,
		Amount:          withdraw.Amount,
		RelAmount:       withdraw.RelAmount,
		BalanceRecordId: withdraw.BalanceRecordId,
		Status:          withdraw.Status,
		Type:            withdraw.Type,
		CreatedAt:       withdraw.CreatedAt,
	}, nil
}

// GetWithdraws .
func (ub *UserBalanceRepo) GetWithdraws(ctx context.Context, b *biz.Pagination, userId int64) ([]*biz.Withdraw, error, int64) {
	var (
		withdraws []*Withdraw
		count     int64
	)
	res := make([]*biz.Withdraw, 0)

	instance := ub.data.db.Table("withdraw")

	if 0 < userId {
		instance = instance.Where("user_id=?", userId)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Order("id desc").Find(&withdraws).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found"), 0
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error()), 0
	}

	for _, withdraw := range withdraws {
		res = append(res, &biz.Withdraw{
			ID:              withdraw.ID,
			UserId:          withdraw.UserId,
			Amount:          withdraw.Amount,
			RelAmount:       withdraw.RelAmount,
			BalanceRecordId: withdraw.BalanceRecordId,
			Status:          withdraw.Status,
			Type:            withdraw.Type,
			CreatedAt:       withdraw.CreatedAt,
		})
	}
	return res, nil, count
}

// GetWithdrawPassOrRewarded .
func (ub *UserBalanceRepo) GetWithdrawPassOrRewarded(ctx context.Context) ([]*biz.Withdraw, error) {
	var withdraws []*Withdraw
	res := make([]*biz.Withdraw, 0)
	if err := ub.data.db.Table("withdraw").Where("status=? or status=?", "pass", "rewarded").Find(&withdraws).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("WITHDRAW_NOT_FOUND", "withdraw not found")
		}

		return nil, errors.New(500, "WITHDRAW ERROR", err.Error())
	}

	for _, withdraw := range withdraws {
		res = append(res, &biz.Withdraw{
			ID:              withdraw.ID,
			UserId:          withdraw.UserId,
			Amount:          withdraw.Amount,
			RelAmount:       withdraw.RelAmount,
			BalanceRecordId: withdraw.BalanceRecordId,
			Status:          withdraw.Status,
			Type:            withdraw.Type,
			CreatedAt:       withdraw.CreatedAt,
		})
	}
	return res, nil
}

// RecommendReward .
func (ub *UserBalanceRepo) RecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "location" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend_vip" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// SystemReward .
func (ub *UserBalanceRepo) SystemReward(ctx context.Context, amount int64, locationId int64) error {
	var (
		reward Reward
		err    error
	)
	reward.UserId = 999999999
	reward.Amount = amount
	reward.BalanceRecordId = 999999999
	reward.Type = "location" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "system_reward" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return err
	}

	return nil
}

// SystemWithdrawReward .
func (ub *UserBalanceRepo) SystemWithdrawReward(ctx context.Context, amount int64, locationId int64) error {
	var (
		reward Reward
		err    error
	)
	reward.UserId = 999999999
	reward.Amount = amount
	reward.BalanceRecordId = 999999999
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "system_reward" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return err
	}

	return nil
}

// GetSystemYesterdayDailyReward .
func (ub *UserBalanceRepo) GetSystemYesterdayDailyReward(ctx context.Context) (*biz.Reward, error) {
	var reward Reward

	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 14 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 14, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 14, 0, 0, 0, time.UTC)

	if err := ub.data.db.
		Where("user_id=?", 999999999).
		Where("created_at>=?", todayStart).
		Where("created_at<?", todayEnd).
		Where("type=?", "system_fee_daily").
		Where("reason=?", "system_fee_daily").
		Table("reward").First(&reward).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}
	return &biz.Reward{
		ID:               reward.ID,
		UserId:           reward.UserId,
		Amount:           reward.Amount,
		BalanceRecordId:  reward.BalanceRecordId,
		Type:             reward.Type,
		TypeRecordId:     reward.TypeRecordId,
		Reason:           reward.Reason,
		ReasonLocationId: reward.ReasonLocationId,
		LocationType:     reward.LocationType,
	}, nil
}

// SystemFee .
func (ub *UserBalanceRepo) SystemFee(ctx context.Context, amount int64, locationId int64) error {
	var (
		reward Reward
		err    error
	)
	reward.UserId = 999999999
	reward.Amount = amount
	reward.BalanceRecordId = 999999999
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "system_fee" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return err
	}

	return nil
}

// UserFee .
func (ub *UserBalanceRepo) UserFee(ctx context.Context, userId int64, amount int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "system_reward" // 本次分红的行为类型
	reward.Reason = "fee"         // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// RecommendWithdrawReward .
func (ub *UserBalanceRepo) RecommendWithdrawReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend_vip" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// NormalRecommendReward .
func (ub *UserBalanceRepo) NormalRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "location" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// NormalWithdrawRecommendReward .
func (ub *UserBalanceRepo) NormalWithdrawRecommendReward(ctx context.Context, userId int64, amount int64, locationId int64) (int64, error) {
	var err error
	if err = ub.data.DB(ctx).Table("user_balance").
		Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance_usdt": gorm.Expr("balance_usdt + ?", amount)}).Error; nil != err {
		return 0, errors.NotFound("user balance err", "user balance not found")
	}

	var userBalance UserBalance
	err = ub.data.DB(ctx).Where(&UserBalance{UserId: userId}).Table("user_balance").First(&userBalance).Error
	if err != nil {
		return 0, err
	}

	var userBalanceRecode UserBalanceRecord
	userBalanceRecode.Balance = userBalance.BalanceUsdt
	userBalanceRecode.UserId = userBalance.UserId
	userBalanceRecode.Type = "reward"
	userBalanceRecode.Amount = amount
	err = ub.data.DB(ctx).Table("user_balance_record").Create(&userBalanceRecode).Error
	if err != nil {
		return 0, err
	}

	var reward Reward
	reward.UserId = userBalance.UserId
	reward.Amount = amount
	reward.BalanceRecordId = userBalanceRecode.ID
	reward.Type = "withdraw" // 本次分红的行为类型
	reward.TypeRecordId = locationId
	reward.Reason = "recommend" // 给我分红的理由
	err = ub.data.DB(ctx).Table("reward").Create(&reward).Error
	if err != nil {
		return 0, err
	}

	return userBalanceRecode.ID, nil
}

// GetUserCurrentMonthRecommendByUserId .
func (uc *UserCurrentMonthRecommendRepo) GetUserCurrentMonthRecommendByUserId(ctx context.Context, userId int64) ([]*biz.UserCurrentMonthRecommend, error) {
	var userCurrentMonthRecommends []*UserCurrentMonthRecommend
	res := make([]*biz.UserCurrentMonthRecommend, 0)
	if err := uc.data.db.Where(&UserCurrentMonthRecommend{UserId: userId}).Table("user_current_month_recommend").Find(&userCurrentMonthRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_CURRENT_MONTH_RECOMMEND_NOT_FOUND", "user current month recommend not found")
		}

		return nil, errors.New(500, "USER CURRENT MONTH RECOMMEND ERROR", err.Error())
	}

	for _, userCurrentMonthRecommend := range userCurrentMonthRecommends {
		res = append(res, &biz.UserCurrentMonthRecommend{
			ID:              userCurrentMonthRecommend.ID,
			UserId:          userCurrentMonthRecommend.UserId,
			RecommendUserId: userCurrentMonthRecommend.RecommendUserId,
			Date:            userCurrentMonthRecommend.Date,
		})
	}
	return res, nil
}

// GetUserCurrentMonthRecommendGroupByUserId .
func (uc *UserCurrentMonthRecommendRepo) GetUserCurrentMonthRecommendGroupByUserId(ctx context.Context, b *biz.Pagination, userId int64) ([]*biz.UserCurrentMonthRecommend, error, int64) {
	var (
		count                      int64
		userCurrentMonthRecommends []*UserCurrentMonthRecommend
	)
	res := make([]*biz.UserCurrentMonthRecommend, 0)

	instance := uc.data.db.Table("user_current_month_recommend")
	if 0 < userId {
		instance = instance.Where("user_id=?", userId)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&userCurrentMonthRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_CURRENT_MONTH_RECOMMEND_NOT_FOUND", "user current month recommend not found"), 0
		}

		return nil, errors.New(500, "USER CURRENT MONTH RECOMMEND ERROR", err.Error()), 0
	}

	for _, userCurrentMonthRecommend := range userCurrentMonthRecommends {
		res = append(res, &biz.UserCurrentMonthRecommend{
			ID:              userCurrentMonthRecommend.ID,
			UserId:          userCurrentMonthRecommend.UserId,
			RecommendUserId: userCurrentMonthRecommend.RecommendUserId,
			Date:            userCurrentMonthRecommend.Date,
		})
	}
	return res, nil, count
}

// GetUserRewardByUserId .
func (ub *UserBalanceRepo) GetUserRewardByUserId(ctx context.Context, userId int64) ([]*biz.Reward, error) {
	var rewards []*Reward
	res := make([]*biz.Reward, 0)
	if err := ub.data.db.Where("user_id", userId).Table("reward").Order("id desc").Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}

	for _, reward := range rewards {
		res = append(res, &biz.Reward{
			ID:               reward.ID,
			UserId:           reward.UserId,
			Amount:           reward.Amount,
			BalanceRecordId:  reward.BalanceRecordId,
			Type:             reward.Type,
			TypeRecordId:     reward.TypeRecordId,
			Reason:           reward.Reason,
			ReasonLocationId: reward.ReasonLocationId,
			LocationType:     reward.LocationType,
			CreatedAt:        reward.CreatedAt,
			AmountB:          reward.AmountB,
		})
	}
	return res, nil
}

// GetUserRewardByUserIds .
func (ub *UserBalanceRepo) GetUserRewardByUserIds(ctx context.Context, userIds ...int64) (map[int64]*biz.UserSortRecommendReward, error) {
	var total []*UserSortRecommendReward
	res := make(map[int64]*biz.UserSortRecommendReward, 0)

	if err := ub.data.db.Table("reward").
		Where("user_id IN(?)", userIds).
		Group("user_id").
		Select("sum(amount) as total, user_id").
		Scopes(Paginate(1, 4)).
		Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}
		return res, errors.New(500, "REWARD ERROR", err.Error())
	}

	for _, v := range total {
		res[v.UserId] = &biz.UserSortRecommendReward{
			Total:  v.Total,
			UserId: v.UserId,
		}
	}

	return res, nil
}

// GetUserRewardTodayTotalByUserId .
func (ub *UserBalanceRepo) GetUserRewardTodayTotalByUserId(ctx context.Context, userId int64) (*biz.UserSortRecommendReward, error) {
	var total *UserSortRecommendReward

	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 14 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}

	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 14, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 14, 0, 0, 0, time.UTC)

	if err := ub.data.db.Table("reward").
		Where("user_id=?", userId).
		Where("created_at>=?", todayStart).Where("created_at<?", todayEnd).
		Select("sum(amount) as total, user_id").
		Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}
		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}

	return &biz.UserSortRecommendReward{
		UserId: total.UserId,
		Total:  total.Total,
	}, nil
}

// GetUserRewards .
func (ub *UserBalanceRepo) GetUserRewards(ctx context.Context, b *biz.Pagination, userId int64) ([]*biz.Reward, error, int64) {
	var (
		rewards []*Reward
		count   int64
	)
	res := make([]*biz.Reward, 0)

	instance := ub.data.db.Table("reward")

	if 0 < userId {
		instance = instance.Where("user_id=?", userId)
	}

	instance = instance.Count(&count)
	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Order("id desc").Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found"), 0
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error()), 0
	}

	for _, reward := range rewards {
		res = append(res, &biz.Reward{
			ID:               reward.ID,
			UserId:           reward.UserId,
			Amount:           reward.Amount,
			BalanceRecordId:  reward.BalanceRecordId,
			Type:             reward.Type,
			TypeRecordId:     reward.TypeRecordId,
			Reason:           reward.Reason,
			ReasonLocationId: reward.ReasonLocationId,
			LocationType:     reward.LocationType,
			CreatedAt:        reward.CreatedAt,
		})
	}
	return res, nil, count
}

// GetUserRewardsLastMonthFee .
func (ub *UserBalanceRepo) GetUserRewardsLastMonthFee(ctx context.Context) ([]*biz.Reward, error) {
	var (
		rewards []*Reward
	)
	res := make([]*biz.Reward, 0)

	instance := ub.data.db.Table("reward")

	now := time.Now().UTC().Add(8 * time.Hour)
	lastMonthFirstDay := now.AddDate(0, -1, -now.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, time.UTC)
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, time.UTC)

	if err := instance.Where("created_at>=?", lastMonthStart).
		Where("created_at<=?", lastMonthEnd).
		Where("reason=?", "system_fee").
		Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error())
	}

	for _, reward := range rewards {
		res = append(res, &biz.Reward{
			ID:               reward.ID,
			UserId:           reward.UserId,
			Amount:           reward.Amount,
			BalanceRecordId:  reward.BalanceRecordId,
			Type:             reward.Type,
			TypeRecordId:     reward.TypeRecordId,
			Reason:           reward.Reason,
			ReasonLocationId: reward.ReasonLocationId,
			LocationType:     reward.LocationType,
			CreatedAt:        reward.CreatedAt,
		})
	}
	return res, nil
}

// CreateUserCurrentMonthRecommend .
func (uc *UserCurrentMonthRecommendRepo) CreateUserCurrentMonthRecommend(ctx context.Context, u *biz.UserCurrentMonthRecommend) (*biz.UserCurrentMonthRecommend, error) {
	var userCurrentMonthRecommend UserCurrentMonthRecommend
	userCurrentMonthRecommend.UserId = u.UserId
	userCurrentMonthRecommend.RecommendUserId = u.RecommendUserId
	userCurrentMonthRecommend.Date = u.Date
	res := uc.data.DB(ctx).Table("user_current_month_recommend").Create(&userCurrentMonthRecommend)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_CURRENT_MONTH_RECOMMEND_ERROR", "用户当月推荐人创建失败")
	}

	return &biz.UserCurrentMonthRecommend{
		ID:              userCurrentMonthRecommend.ID,
		UserId:          userCurrentMonthRecommend.UserId,
		RecommendUserId: userCurrentMonthRecommend.RecommendUserId,
		Date:            userCurrentMonthRecommend.Date,
	}, nil
}

// GetUserBalanceByUserIds .
func (ub UserBalanceRepo) GetUserBalanceByUserIds(ctx context.Context, userIds ...int64) (map[int64]*biz.UserBalance, error) {
	var userBalances []*UserBalance
	res := make(map[int64]*biz.UserBalance)
	if err := ub.data.db.Where("user_id IN (?)", userIds).Table("user_balance").Find(&userBalances).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return nil, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	for _, userBalance := range userBalances {
		res[userBalance.UserId] = &biz.UserBalance{
			ID:          userBalance.ID,
			UserId:      userBalance.UserId,
			BalanceUsdt: userBalance.BalanceUsdt,
			BalanceDhb:  userBalance.BalanceDhb,
		}
	}

	return res, nil
}

type UserBalanceTotal struct {
	Total int64
}

type UserSortRecommendReward struct {
	UserId int64
	Total  int64
}

// GetUserBalanceUsdtTotal .
func (ub UserBalanceRepo) GetUserBalanceUsdtTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance").Select("sum(balance_usdt) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserBalanceRecordUsdtTotal .
func (ub UserBalanceRepo) GetUserBalanceRecordUsdtTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "deposit").
		Where("coin_type=?", "usdt").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserBalanceRecordUserUsdtTotal .
func (ub UserBalanceRepo) GetUserBalanceRecordUserUsdtTotal(ctx context.Context, userId int64) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_record").
		Where("user_id=?", userId).
		Where("type=?", "deposit").
		Where("coin_type=?", "usdt").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserBalanceRecordUsdtTotalToday .
func (ub UserBalanceRepo) GetUserBalanceRecordUsdtTotalToday(ctx context.Context) (int64, error) {
	var total UserBalanceTotal

	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 14 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 15, 59, 59, 0, time.UTC)

	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "deposit").
		Where("coin_type=?", "usdt").
		Where("created_at>=?", todayStart).Where("created_at<=?", todayEnd).
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserRewardRecommendSort .
func (ub *UserBalanceRepo) GetUserRewardRecommendSort(ctx context.Context) ([]*biz.UserSortRecommendReward, error) {
	var total []*UserSortRecommendReward
	res := make([]*biz.UserSortRecommendReward, 0)

	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 14 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 14, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 14, 0, 0, 0, time.UTC)

	if err := ub.data.db.Table("reward").
		Where("type=?", "location").
		Where("reason=?", "recommend").
		Where("created_at>=?", todayStart).
		Where("created_at<?", todayEnd).
		Group("user_id").
		Select("sum(amount) as total, user_id").
		Order("total desc").
		Scopes(Paginate(1, 4)).
		Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found")
		}
		return res, errors.New(500, "REWARD ERROR", err.Error())
	}

	for _, v := range total {
		res = append(res, &biz.UserSortRecommendReward{
			Total:  v.Total,
			UserId: v.UserId,
		})
	}

	return res, nil
}

// GetUserWithdrawUsdtTotalToday .
func (ub UserBalanceRepo) GetUserWithdrawUsdtTotalToday(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	now := time.Now().UTC()
	var startDate time.Time
	var endDate time.Time
	if 14 <= now.Hour() {
		startDate = now
		endDate = now.AddDate(0, 0, 1)
	} else {
		startDate = now.AddDate(0, 0, -1)
		endDate = now
	}
	todayStart := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 16, 0, 0, 0, time.UTC)
	todayEnd := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 15, 59, 59, 0, time.UTC)
	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "withdraw").
		Where("coin_type=?", "usdt").
		Where("created_at>=?", todayStart).Where("created_at<=?", todayEnd).
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserWithdrawUsdtTotal .
func (ub UserBalanceRepo) GetUserWithdrawUsdtTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "withdraw").
		Where("coin_type=?", "usdt").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserRewardUsdtTotal .
func (ub UserBalanceRepo) GetUserRewardUsdtTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("user_balance_record").
		Where("type=?", "reward").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetSystemRewardUsdtTotal .
func (ub UserBalanceRepo) GetSystemRewardUsdtTotal(ctx context.Context) (int64, error) {
	var total UserBalanceTotal
	if err := ub.data.db.Table("reward").
		Where("reason=? or reason=?", "system_reward", "system_fee").
		Select("sum(amount) as total").Take(&total).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return total.Total, errors.NotFound("USER_BALANCE_RECORD_NOT_FOUND", "user balance not found")
		}

		return total.Total, errors.New(500, "USER BALANCE RECORD ERROR", err.Error())
	}

	return total.Total, nil
}

// GetUserInfoByUserIds .
func (ui *UserInfoRepo) GetUserInfoByUserIds(ctx context.Context, userIds ...int64) (map[int64]*biz.UserInfo, error) {
	var userInfos []*UserInfo
	res := make(map[int64]*biz.UserInfo, 0)
	if err := ui.data.db.Where("user_id IN (?)", userIds).Table("user_info").Find(&userInfos).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USERINFO_NOT_FOUND", "userinfo not found")
		}

		return nil, errors.New(500, "USERINFO ERROR", err.Error())
	}

	for _, userInfo := range userInfos {
		res[userInfo.UserId] = &biz.UserInfo{
			ID:               userInfo.ID,
			UserId:           userInfo.UserId,
			Vip:              userInfo.Vip,
			HistoryRecommend: userInfo.HistoryRecommend,
			TeamCsdBalance:   userInfo.TeamCsdBalance,
		}
	}

	return res, nil
}

// GetUserCurrentMonthRecommendCountByUserIds .
func (uc *UserCurrentMonthRecommendRepo) GetUserCurrentMonthRecommendCountByUserIds(ctx context.Context, userIds ...int64) (map[int64]int64, error) {
	var userCurrentMonthRecommends []*UserCurrentMonthRecommend
	res := make(map[int64]int64, 0)
	if err := uc.data.db.Where("user_id IN (?)", userIds).Table("user_current_month_recommend").Find(&userCurrentMonthRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_CURRENT_MONTH_RECOMMEND_NOT_FOUND", "user current month recommend not found")
		}

		return nil, errors.New(500, "USER CURRENT MONTH RECOMMEND ERROR", err.Error())
	}

	for _, userCurrentMonthRecommend := range userCurrentMonthRecommends {
		if _, ok := res[userCurrentMonthRecommend.UserId]; !ok {
			res[userCurrentMonthRecommend.UserId] = 1
		} else {
			res[userCurrentMonthRecommend.UserId]++
		}
	}
	return res, nil
}

// GetUserLastMonthRecommend .
func (uc *UserCurrentMonthRecommendRepo) GetUserLastMonthRecommend(ctx context.Context) ([]int64, error) {
	var userCurrentMonthRecommends []*UserCurrentMonthRecommend
	res := make([]int64, 0)

	now := time.Now().UTC().Add(8 * time.Hour)
	lastMonthFirstDay := now.AddDate(0, -1, -now.Day()+1)
	lastMonthStart := time.Date(lastMonthFirstDay.Year(), lastMonthFirstDay.Month(), lastMonthFirstDay.Day(), 0, 0, 0, 0, time.UTC)
	lastMonthEndDay := lastMonthFirstDay.AddDate(0, 1, -1)
	lastMonthEnd := time.Date(lastMonthEndDay.Year(), lastMonthEndDay.Month(), lastMonthEndDay.Day(), 23, 59, 59, 0, time.UTC)

	if err := uc.data.db.Table("user_current_month_recommend").
		Group("user_id").
		Having("count(id) >= 5").
		Where("date>=?", lastMonthStart).
		Where("date<=?", lastMonthEnd).
		Find(&userCurrentMonthRecommends).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("USER_CURRENT_MONTH_RECOMMEND_NOT_FOUND", "user current month recommend not found")
		}

		return nil, errors.New(500, "USER CURRENT MONTH RECOMMEND ERROR", err.Error())
	}

	for _, userCurrentMonthRecommend := range userCurrentMonthRecommends {
		res = append(res, userCurrentMonthRecommend.UserId)
	}
	return res, nil
}
