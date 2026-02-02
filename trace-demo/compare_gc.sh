#!/bin/bash

# GC 优化效果对比脚本

echo "🔬 GC 优化效果对比"
echo "=================================="
echo ""
echo "测试配置: 50000 任务，每任务 ~100KB 内存分配"
echo ""

# 创建输出目录
mkdir -p traces

echo "1️⃣  Worker Pool (GOGC=100, 默认)"
echo "-----------------------------------"
go run . -mode=pool -tasks=50000
echo ""

echo "2️⃣  Worker Pool (GOGC=1000, 优化)"
echo "-----------------------------------"
go run . -mode=pool-tuned -tasks=50000 -gogc=1000
echo ""

echo "=================================="
echo "📊 生成对比 trace 文件..."
echo "=================================="
echo ""

echo "生成 GOGC=100 trace..."
go run . -mode=pool -tasks=50000 -trace=traces/gc100.trace > /dev/null 2>&1
echo "✅ traces/gc100.trace"

echo "生成 GOGC=1000 trace..."
go run . -mode=pool-tuned -tasks=50000 -trace=traces/gc1000.trace -gogc=1000 > /dev/null 2>&1
echo "✅ traces/gc1000.trace"

echo ""
echo "=================================="
echo "✅ 对比完成！"
echo "=================================="
echo ""
echo "📖 查看 trace 对比："
echo "   go tool trace traces/gc100.trace"
echo "   go tool trace traces/gc1000.trace"
echo ""
echo "💡 在 trace 中观察："
echo "   - GC 事件频率（GOGC=100 更频繁）"
echo "   - GC 暂停时间"
echo "   - Goroutine 执行情况"
echo ""
echo "📚 查看详细分析："
echo "   cat GC_OPTIMIZATION.md"
echo ""
