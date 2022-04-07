package ntpTime

import (
	"github.com/beevik/ntp"
	"time"
)

const (
	host = "0.beevik-ntp.pool.ntp.org"
)

func TimeWithTimeZone(t time.Time) (time.Time, error) {
	r, err := ntp.Query(host)
	if err != nil {
		return time.Time{}, err
	}
	// Validate метод выполняет дополнительные проверки работоспособности, чтобы определить, подходит ли ответ для целей синхронизации времени
	err = r.Validate()
	if err == nil {
		// ClockOffset: смещение часов локальной системы относительно часов сервера (+3 MSK)
		t = t.Add(r.ClockOffset)

		return t, nil
	}

	return time.Time{}, err
}
