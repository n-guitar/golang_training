export GOMAXPROCS=1
echo "GOMAXPROCS=${GOMAXPROCS}"
time go run ch08_ex05_goroutine_mandelbort.go > /dev/null 2>&1

echo ""
export GOMAXPROCS=2
echo "GOMAXPROCS=${GOMAXPROCS}"
time go run ch08_ex05_goroutine_mandelbort.go > /dev/null 2>&1

echo ""
export GOMAXPROCS=8
echo "GOMAXPROCS=${GOMAXPROCS}"
time go run ch08_ex05_goroutine_mandelbort.go > /dev/null 2>&1

echo ""
export GOMAXPROCS=16
echo "GOMAXPROCS=${GOMAXPROCS}"
time go run ch08_ex05_goroutine_mandelbort.go > /dev/null 2>&1

