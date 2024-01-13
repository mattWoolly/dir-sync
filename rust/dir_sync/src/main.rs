use std::{collections::HashSet, env, fs, io, path::{Path, PathBuf}};

fn main() -> io::Result<()> {
    let args: Vec<String> = env::args().collect();
    if args.len() != 3 { return Err(io::Error::new(io::ErrorKind::InvalidInput, "Usage: sync_dirs <dir1> <dir2>")); }
    
    let dir1 = &args[1];
    let dir2 = &args[2];
    sync_directories(Path::new(dir1), Path::new(dir2))?;
    Ok(())
}

fn sync_directories(dir1: &Path, dir2: &Path) -> io::Result<()> {
    let entries_dir1 = read_directory(dir1, dir1)?;
    let entries_dir2 = read_directory(dir2, dir2)?;
    let unique_in_dir1: HashSet<_> = entries_dir1.difference(&entries_dir2).collect();
    let unique_in_dir2: HashSet<_> = entries_dir2.difference(&entries_dir1).collect();

    for file in unique_in_dir1.union(&unique_in_dir2) {
        let (src, dest) = if entries_dir1.contains(file) { (dir1.join(file), dir2.join(file)) } else { (dir2.join(file), dir1.join(file)) };
        if let Some(parent) = dest.parent() { fs::create_dir_all(parent)?; }
        fs::copy(&src, &dest)?;
    }
    Ok(())
}

fn read_directory(base: &Path, dir: &Path) -> io::Result<HashSet<PathBuf>> {
    let mut entries = HashSet::new();
    for entry in fs::read_dir(dir)? {
        let entry = entry?;
        let path = entry.path();
        if path.is_dir() { entries.extend(read_directory(base, &path)?); } else { entries.insert(path.strip_prefix(base)?.to_path_buf()); }
    }
    Ok(entries)
}
