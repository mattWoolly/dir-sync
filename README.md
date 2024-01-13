# Dir-Sync

Dir-Sync is a utility for synchronizing the contents of two directories. It ensures that both directories have the same set of files by copying unique files from one directory to the other. This tool is available in Python, Rust, and now Go (Golang), offering flexibility depending on the user's environment or preferences and for fun comparing performance.

## Features

- **Cross-Language Support**: Available in Python, Rust, and Go.
- **Directory Synchronization**: Identifies and syncs unique files between two directories.
- **Recursive Traversal**: Capable of traversing through subdirectories to ensure complete synchronization.

## Usage

### Python Version

Located in the `python/` directory. To use the Python version of Dir-Sync, run the following command:

```bash
python main.py <dir1> <dir2>
Replace <dir1> and <dir2> with the paths of the directories you want to synchronize.
```

### Rust Version
Located in the rust/dir_sync/ directory. To use the Rust version, first compile the program using Cargo, and then run the compiled binary:

```bash
Copy code
cd rust/dir_sync/
cargo build --release
./target/release/dir_sync <dir1> <dir2>
Replace <dir1> and <dir2> with the paths of the directories you want to synchronize.
```

## Go Version
Located in the go/ directory. To use the Go version of Dir-Sync, first compile the program using go build, and then run the compiled binary:

```bash
Copy code
cd go/
go build -o dir_sync
./dir_sync <dir1> <dir2>
```
Again, replace <dir1> and <dir2> with the paths of the directories you want to synchronize.

### Requirements
Python 3 for the Python version.
Rust toolchain for the Rust version.
Go compiler for the Go version.
Contributing
Contributions to Dir-Sync are welcome. Please feel free to submit pull requests or open issues to suggest improvements or report bugs.

## License

This project is open-source and available under the [MIT License](#mit-license).

---

## MIT License

MIT License

Copyright (c) 2024 Matt Woolly

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

**The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.**

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. **IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.**