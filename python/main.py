import os
import shutil
import sys
from pathlib import Path

def main():
    try:
        dir1, dir2 = get_directories_from_args()
        sync_directories(dir1, dir2)
    except ValueError as e:
        print(e, file=sys.stderr)
        sys.exit(1)

def get_directories_from_args():
    if len(sys.argv) != 3:
        raise ValueError("Usage: sync_dirs <dir1> <dir2>")
    return Path(sys.argv[1]), Path(sys.argv[2])

def sync_directories(dir1, dir2):
    entries_dir1 = set(read_directory(dir1))
    entries_dir2 = set(read_directory(dir2))

    sync_unique_files(entries_dir1 - entries_dir2, dir1, dir2)
    sync_unique_files(entries_dir2 - entries_dir1, dir2, dir1)

def sync_unique_files(unique_files, src_dir, dest_dir):
    for file in unique_files:
        src_path = src_dir / file
        dest_path = dest_dir / file
        dest_path.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(src_path, dest_path)

def read_directory(dir_path):
    return [path.relative_to(dir_path) for path in dir_path.rglob('*') if path.is_file()]

if __name__ == "__main__":
    main()
