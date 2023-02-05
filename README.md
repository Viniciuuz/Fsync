# Fsync

CLI app for automating folder backup

## Win x64 compiled binaries [here](https://github.com/Viniciuuz/Fsync/releases) 

## Usage

### config
```ps
PS C:\Go\Fsync> fsync config
```

the *config* flag will open the "Select folder" for the source and the destination, then it will ask if the paths selected are correct, if you click yes, the path will be saved to "config.ini"


---

### list

```ps
PS C:\Go\Fsync> fsync list
```

the *list* flag will list all folders "queued" to sync

Example:

```ps
PS C:\Go\Fsync> fsync list

Total Paths: 2

src: G:\test
dest: G:\test2

src: G:\text
dest: C:\Users\user123\Documents\test
```

---

### sync

```ps
PS C:\Go\Fsync> fsync sync
```

the *sync* flag will sync all the folders registered in "config.ini"

Example:
```ps
PS C:\Go\Fsync> fsync sync

Backing up: G:\test
G:\test2\copy\readme.txt already exists
copying G:\test\index.html
G:\test2\readme - Copy.txt already exists
G:\test2\readme.txt already exists

Backing up: G:\text
copying C:\Users\user123\Documents\test\text.txt