killall clock

TZ=US/Eastern go run clock.go -p 8010 &
TZ=Asia/Tokyo go run clock.go -p 8020 &
TZ=Europe/London go run clock.go -p 8030 &


go run clockwall.go NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030


