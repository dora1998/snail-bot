package commands

import (
	"github.com/dora1998/snail-bot/mock_repository"
	"github.com/dora1998/snail-bot/mock_twitter"
	"github.com/dora1998/snail-bot/repository"
	"github.com/dora1998/snail-bot/testutil"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestCommandHandler_add(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	fakeTime := time.Date(2020, 1, 1, 0, 0, 0, 0, loc)
	testutil.SetFakeTime(fakeTime)

	type args struct {
		body     string
		username string
		statusId int64
	}

	tests := []struct {
		name     string
		args     args
		injector func(c *gomock.Controller) (*mock_repository.MockRepository, *mock_twitter.MockTwitterClient)
	}{
		{
			"normal",
			args{
				body:     "test 12/31",
				username: "testuser",
				statusId: 0,
			},
			func(c *gomock.Controller) (*mock_repository.MockRepository, *mock_twitter.MockTwitterClient) {
				deadline := time.Date(2020, 12, 31, 0, 0, 0, 0, loc)

				mr := mock_repository.NewMockRepository(c)
				mr.EXPECT().Add(gomock.Eq("test"), gomock.Eq(deadline), gomock.Eq("testuser")).Return(&repository.Task{
					Id:        "hoge",
					Body:      "test",
					Deadline:  deadline,
					CreatedAt: time.Time{},
					CreatedBy: "testuser",
				}).Times(1)

				mt := mock_twitter.NewMockTwitterClient(c)
				mt.EXPECT().IsFollowing(gomock.Eq("testuser")).Return(true).Times(1)
				mt.EXPECT().Reply(gomock.Eq("タスクを追加しました！\ntest (2020/12/31)"), gomock.Eq(int64(0))).Times(1)
				return mr, mt
			},
		},
		{
			"not following",
			args{
				body:     "test 12/31",
				username: "testuser",
				statusId: 0,
			},
			func(c *gomock.Controller) (*mock_repository.MockRepository, *mock_twitter.MockTwitterClient) {
				mt := mock_twitter.NewMockTwitterClient(c)
				mt.EXPECT().IsFollowing(gomock.Eq("testuser")).Return(false).Times(1)
				mt.EXPECT().Reply(gomock.Eq("この操作はフォローされている人しかできません🙇‍♂️"), gomock.Eq(int64(0))).Times(1)
				return nil, mt
			},
		},
		{
			"incorrect format",
			args{
				body:     "test",
				username: "testuser",
				statusId: 0,
			},
			func(c *gomock.Controller) (*mock_repository.MockRepository, *mock_twitter.MockTwitterClient) {
				mt := mock_twitter.NewMockTwitterClient(c)
				mt.EXPECT().IsFollowing(gomock.Eq("testuser")).Return(true).Times(1)
				mt.EXPECT().Reply(gomock.Eq("タスクの追加に失敗しました…"), gomock.Eq(int64(0))).Times(1)
				return nil, mt
			},
		},
		{
			"incorrect date format",
			args{
				body:     "test 15/99",
				username: "testuser",
				statusId: 0,
			},
			func(c *gomock.Controller) (*mock_repository.MockRepository, *mock_twitter.MockTwitterClient) {
				mt := mock_twitter.NewMockTwitterClient(c)
				mt.EXPECT().IsFollowing(gomock.Eq("testuser")).Return(true).Times(1)
				mt.EXPECT().Reply(gomock.Eq("タスクの追加に失敗しました…"), gomock.Eq(int64(0))).Times(1)
				return nil, mt
			},
		},
		{
			"db error",
			args{
				body:     "test 12/31",
				username: "testuser",
				statusId: 0,
			},
			func(c *gomock.Controller) (*mock_repository.MockRepository, *mock_twitter.MockTwitterClient) {
				deadline := time.Date(2020, 12, 31, 0, 0, 0, 0, loc)

				mr := mock_repository.NewMockRepository(c)
				mr.EXPECT().Add(gomock.Eq("test"), gomock.Eq(deadline), gomock.Eq("testuser")).Return(nil).Times(1)

				mt := mock_twitter.NewMockTwitterClient(c)
				mt.EXPECT().IsFollowing(gomock.Eq("testuser")).Return(true).Times(1)
				mt.EXPECT().Reply(gomock.Eq("タスクの追加に失敗しました…"), gomock.Eq(int64(0))).Times(1)
				return mr, mt
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo, twitterClient := tt.injector(ctrl)
			h := &CommandHandler{
				repository:    repo,
				twitterClient: twitterClient,
			}
			h.add(tt.args.body, tt.args.username, tt.args.statusId)
		})
	}
}
