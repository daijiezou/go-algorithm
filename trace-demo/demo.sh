#!/bin/bash

# 快速演示脚本

echo "🚀 Go Trace 性能分析演示"
echo "=================================="
echo ""

# 创建输出目录
mkdir -p traces

echo "📊 1. 运行基础性能对比（50000 任务，确保触发 GC）..."
echo ""
go run . -mode=all -tasks=50000
echo ""

echo "📈 2. 生成 trace 文件..."
echo ""
echo "   生成 Fan-out trace..."
go run . -mode=fanout -tasks=50000 -trace=traces/fanout.trace > /dev/null
echo "   ✅ traces/fanout.trace"

echo "   生成 Worker Pool trace..."
go run . -mode=pool -tasks=50000 -trace=traces/pool.trace > /dev/null
echo "   ✅ traces/pool.trace"

echo "   生成 Worker Pool (优化) trace..."
go run . -mode=pool-tuned -tasks=50000 -trace=traces/pool-tuned.trace -gogc=1000 > /dev/null
echo "   ✅ traces/pool-tuned.trace"

echo ""
echo "🔬 3. 运行 Benchmark 测试..."
echo ""
go test -bench='BenchmarkSingleThreaded|BenchmarkFanOut|BenchmarkWorkerPool' -benchmem -benchtime=1s

echo ""
echo "⚡ 4. False Sharing 性能对比..."
echo ""
go test -bench='BenchmarkFalseSharing|BenchmarkPaddedCounters' -benchmem -benchtime=1s

echo ""
echo "=================================="
echo "✅ 演示完成！"
echo "=================================="
echo ""
echo "📖 查看 trace 文件："
echo "   go tool trace traces/fanout.trace"
echo "   go tool trace traces/pool.trace"
echo "   go tool trace traces/pool-tuned.trace"
echo ""
echo "📚 查看完整文档："
echo "   cat README.md"
echo "   cat USAGE.md"
echo ""
echo "🧪 运行完整测试："
echo "   ./run.sh"
echo ""
