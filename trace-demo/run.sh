#!/bin/bash

# Go Trace 演示脚本

echo "==================================="
echo "Go Trace 性能分析演示"
echo "==================================="

# 创建输出目录
mkdir -p traces

echo ""
echo "1️⃣  运行所有模式并生成 trace..."
go run . -mode=all -tasks=10000 -trace=traces/all.trace

echo ""
echo "2️⃣  运行单线程模式..."
go run . -mode=single -tasks=10000 -trace=traces/single.trace

echo ""
echo "3️⃣  运行 Fan-out 模式..."
go run . -mode=fanout -tasks=10000 -trace=traces/fanout.trace

echo ""
echo "4️⃣  运行 Worker Pool 模式（默认 GOGC=100）..."
go run . -mode=pool -tasks=10000 -trace=traces/pool.trace

echo ""
echo "5️⃣  运行 Worker Pool 模式（优化 GOGC=1000）..."
go run . -mode=pool-tuned -tasks=10000 -trace=traces/pool-tuned.trace -gogc=1000

echo ""
echo "==================================="
echo "运行 Benchmark 测试..."
echo "==================================="

echo ""
echo "📊 基础 Benchmark..."
go test -bench=. -benchmem -benchtime=3s

echo ""
echo "📊 使用不同 GOGC 值的 Benchmark..."
GOGC=100 go test -bench=BenchmarkWorkerPool -benchmem -benchtime=3s
GOGC=1000 go test -bench=BenchmarkWorkerPool -benchmem -benchtime=3s

echo ""
echo "==================================="
echo "检测数据竞争..."
echo "==================================="
echo ""
echo "⚠️  运行数据竞争检测（这会很慢）..."
go run -race . -mode=single -tasks=100

echo ""
echo "==================================="
echo "生成性能分析文件..."
echo "==================================="

echo ""
echo "📈 CPU Profile..."
go test -bench=BenchmarkWorkerPool -cpuprofile=traces/cpu.prof -benchtime=5s

echo ""
echo "📈 Memory Profile..."
go test -bench=BenchmarkWorkerPool -memprofile=traces/mem.prof -benchtime=5s

echo ""
echo "==================================="
echo "✅ 完成！"
echo "==================================="
echo ""
echo "查看 trace 文件："
echo "  go tool trace traces/all.trace"
echo "  go tool trace traces/fanout.trace"
echo "  go tool trace traces/pool.trace"
echo "  go tool trace traces/pool-tuned.trace"
echo ""
echo "查看 CPU profile："
echo "  go tool pprof traces/cpu.prof"
echo ""
echo "查看 Memory profile："
echo "  go tool pprof traces/mem.prof"
echo ""
echo "对比 False Sharing 性能："
echo "  go test -bench='BenchmarkFalseSharing|BenchmarkPaddedCounters' -benchmem"
echo ""
