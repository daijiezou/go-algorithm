#!/usr/bin/env python3
"""
WebM 转 MP4 转换工具
使用 ffmpeg-python 库进行视频格式转换
"""

import os
import sys
import argparse
from pathlib import Path

try:
    import ffmpeg
except ImportError:
    print("错误: 需要安装 ffmpeg-python 库")
    print("请运行: pip install ffmpeg-python")
    print("同时确保系统已安装 ffmpeg: brew install ffmpeg (macOS)")
    sys.exit(1)


def convert_webm_to_mp4(input_file, output_file=None, overwrite=False):
    """
    将 WebM 文件转换为 MP4 格式
    
    Args:
        input_file: 输入的 WebM 文件路径
        output_file: 输出的 MP4 文件路径 (可选，默认为同名 .mp4 文件)
        overwrite: 是否覆盖已存在的文件
    
    Returns:
        bool: 转换是否成功
    """
    input_path = Path(input_file)
    
    # 检查输入文件是否存在
    if not input_path.exists():
        print(f"错误: 输入文件不存在: {input_file}")
        return False
    
    # 如果没有指定输出文件，使用同名的 .mp4 文件
    if output_file is None:
        output_file = input_path.with_suffix('.mp4')
    else:
        output_file = Path(output_file)
    
    # 检查输出文件是否已存在
    if output_file.exists() and not overwrite:
        print(f"错误: 输出文件已存在: {output_file}")
        print("使用 --overwrite 参数强制覆盖")
        return False
    
    try:
        print(f"开始转换: {input_path.name} -> {output_file.name}")
        
        # 使用 ffmpeg 进行转换
        # 使用 libopenh264 编码器（你的 ffmpeg 支持的编码器）
        stream = ffmpeg.input(str(input_path))
        stream = ffmpeg.output(
            stream,
            str(output_file),
            vcodec='libopenh264',  # 使用 OpenH264 编码
            acodec='aac',          # 使用 AAC 音频编码
            video_bitrate='5M',    # 视频比特率
            audio_bitrate='192k'   # 音频比特率
        )
        
        # 如果需要覆盖，添加 overwrite_output 参数
        if overwrite:
            stream = ffmpeg.overwrite_output(stream)
        
        # 执行转换
        ffmpeg.run(stream, capture_stdout=True, capture_stderr=True)
        
        print(f"✓ 转换成功: {output_file}")
        print(f"  输入文件大小: {input_path.stat().st_size / 1024 / 1024:.2f} MB")
        print(f"  输出文件大小: {output_file.stat().st_size / 1024 / 1024:.2f} MB")
        return True
        
    except ffmpeg.Error as e:
        print(f"错误: 转换失败")
        print(f"错误信息: {e.stderr.decode() if e.stderr else str(e)}")
        return False
    except Exception as e:
        print(f"错误: {str(e)}")
        return False


def convert_directory(input_dir, output_dir=None, overwrite=False):
    """
    批量转换目录中的所有 WebM 文件
    
    Args:
        input_dir: 输入目录
        output_dir: 输出目录 (可选，默认为输入目录)
        overwrite: 是否覆盖已存在的文件
    """
    input_path = Path(input_dir)
    
    if not input_path.is_dir():
        print(f"错误: {input_dir} 不是一个有效的目录")
        return
    
    # 查找所有 .webm 文件
    webm_files = list(input_path.glob('*.webm'))
    
    if not webm_files:
        print(f"在 {input_dir} 中没有找到 .webm 文件")
        return
    
    print(f"找到 {len(webm_files)} 个 WebM 文件")
    
    # 如果指定了输出目录，确保它存在
    if output_dir:
        output_path = Path(output_dir)
        output_path.mkdir(parents=True, exist_ok=True)
    else:
        output_path = input_path
    
    # 转换每个文件
    success_count = 0
    for webm_file in webm_files:
        output_file = output_path / webm_file.with_suffix('.mp4').name
        if convert_webm_to_mp4(webm_file, output_file, overwrite):
            success_count += 1
        print()  # 空行分隔
    
    print(f"完成: 成功转换 {success_count}/{len(webm_files)} 个文件")


def main():
    parser = argparse.ArgumentParser(
        description='将 WebM 文件转换为 MP4 格式',
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
示例:
  # 转换单个文件
  python webm_to_mp4.py input.webm
  
  # 指定输出文件名
  python webm_to_mp4.py input.webm -o output.mp4
  
  # 批量转换目录中的所有文件
  python webm_to_mp4.py -d /path/to/directory
  
  # 覆盖已存在的文件
  python webm_to_mp4.py input.webm --overwrite
        """
    )
    
    parser.add_argument('input', nargs='?', help='输入的 WebM 文件路径')
    parser.add_argument('-o', '--output', help='输出的 MP4 文件路径')
    parser.add_argument('-d', '--directory', help='批量转换目录中的所有 WebM 文件')
    parser.add_argument('--overwrite', action='store_true', help='覆盖已存在的输出文件')
    
    args = parser.parse_args()
    
    # 检查是否提供了输入
    if not args.input and not args.directory:
        parser.print_help()
        sys.exit(1)
    
    # 批量转换模式
    if args.directory:
        convert_directory(args.directory, args.output, args.overwrite)
    # 单文件转换模式
    elif args.input:
        success = convert_webm_to_mp4(args.input, args.output, args.overwrite)
        sys.exit(0 if success else 1)


if __name__ == '__main__':
    main()
