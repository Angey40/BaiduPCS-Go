package uploader

import (
	"time"
)

// Status 上传状态
type Status interface {
	TotalSize() int64           // 总大小
	Uploaded() int64            // 已上传数据
	SpeedsPerSecond() int64     // 每秒的上传速度
	TimeElapsed() time.Duration // 上传时间
}

// UploadStatus 上传状态
type UploadStatus struct {
	totalSize       int64         // 总大小
	uploaded        int64         // 已上传数据
	speedsPerSecond int64         // 每秒的上传速度
	timeElapsed     time.Duration // 上传时间
}

// TotalSize 返回总大小
func (us *UploadStatus) TotalSize() int64 {
	return us.totalSize
}

// Uploaded 返回已上传数据
func (us *UploadStatus) Uploaded() int64 {
	return us.uploaded
}

// SpeedsPerSecond 返回每秒的上传速度
func (us *UploadStatus) SpeedsPerSecond() int64 {
	return us.speedsPerSecond
}

// TimeElapsed 返回上传时间
func (us *UploadStatus) TimeElapsed() time.Duration {
	return us.timeElapsed
}

// GetStatusChan 获取上传状态
func (u *Uploader) GetStatusChan() <-chan Status {
	c := make(chan Status)

	go func() {
		for {
			select {
			case <-u.finished:
				close(c)
				return
			default:
				if !u.executed {
					time.Sleep(1 * time.Second)
					continue
				}

				old := u.readed64.Readed()
				time.Sleep(1 * time.Second) // 每秒统计

				readed := u.readed64.Readed()
				c <- &UploadStatus{
					totalSize:       u.readed64.Len(),
					uploaded:        readed,
					speedsPerSecond: readed - old,
					timeElapsed:     time.Since(u.executeTime) / 1000000 * 1000000,
				}
			}
		}
	}()
	return c
}

// GetStatusChan 获取上传状态
func (muer *MultiUploader) GetStatusChan() <-chan Status {
	muer.lazyInit()
	c := make(chan Status)

	go func() {
		for {
			select {
			case <-muer.finished:
				close(c)
				return
			default:
				if !muer.executed {
					time.Sleep(1 * time.Second)
					continue
				}

				old := muer.workers.Readed()
				time.Sleep(1 * time.Second) // 每秒统计

				readed := muer.workers.Readed()
				c <- &UploadStatus{
					totalSize:       muer.file.Len(),
					uploaded:        readed,
					speedsPerSecond: readed - old,
					timeElapsed:     time.Since(muer.executeTime) / 1000000 * 1000000,
				}
			}
		}
	}()
	return c
}

// UpdateInstanceStateChan 更新状态的信号
func (muer *MultiUploader) UpdateInstanceStateChan() <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case signal := <-muer.updateInstanceStateChan:
				c <- signal
			case <-muer.finished:
				return
			}
		}
	}()
	return c
}
