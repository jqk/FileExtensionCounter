# FileExtensionCounter

FileExtensionCounter is a command line program used to search all file extensions under a given path, and give the number of files and total file size for each extension.

Only tested on Windows but should also run on Linux and Mac operating systems.

FileExtensionCounter 是一个命令行程序，用于搜索给定路径下的所有文件扩展名，并给出每个扩展名的文件数量和总文件大小。该程序仅在 Windows 上测试过，但也应该可以在 Linux 和 Mac 操作系统上运行。

## Usage

```text {.line-numbers}
$ fec

Copyright (c) 1999-2023 Not a dream Co., Ltd.
file extension counter (fec) 1.2.0, 2023-09-11

Argument error: wrong number of argument

Usage:
  fec [<byPassPermissionError> <caseSensitive> <recursive> <sortMethod>] <path/to/counting/extensions>
       counting extensions in specified path

Options:
  All options must be either specified in order or omitted. Default values are used when omitted: -t -f -t -e

  byPassPermissionError: -t is true, skip the permission error (default); -f is false, throw an error
  caseSensitive        : -t is true, the extension is case sensitive; -f is false, the extension is case insensitive (default)
  recursive            : -t is true, sub directories are included (default); -f is false, sub directories are excluded
  sortMethod           : how to sort the result
                         -c means sort by count
                         -e means sort by extension (default)
                         -s means sort by size

Otherwise: show this help.
See <https://github.com/jqk/FileExtensionCounter> for more information.
```

## Example

### Search files extensions in case insensitive manner

```bash {.line-numbers}
$ fec e:\temp
Or
$ fec -t -f -t -e e:\temp

Copyright (c) 1999-2023 Not a dream Co., Ltd.
file extension counter (fec) 1.2.0, 2023-09-11

Searching...
Searching done.

Searching path          : e:\temp
Bypass permission error : true
Case sensitive          : false
Recursive               : true
Found file              : 3126
Found extension         : 78
Total size              : 5.063 GB
Elapsed time            : 69.9484ms

  No.  Extension        Count          Size

    1                     537    696.310 KB
    2  ._                  54       0 bytes
    3  .apk                 1     31.922 MB
    4  .at                 76    363.312 KB
    5  .bat                 7     15.746 KB
   ....
   69  .toml                4     711 bytes
   70  .txt               322      1.775 MB
   71  .values              5      1.582 MB
   72  .vsix                1     56.140 KB
   73  .woff2               9   1004.168 KB
   74  .xaml                2     929 bytes
   75  .xls                 2    719.000 KB
   76  .xlsx                1     11.426 KB
   77  .xml               139     27.135 MB
   78  .zip                13      4.392 GB
```

### Search files extensions in case sensitive manner

Files with uppercase extension `.TXT` are treated differently from files with lowercase `.txt`.

```bash {.line-numbers}
$ fec -t -t -t -e e:/temp

Copyright (c) 1999-2023 Not a dream Co., Ltd.
file extension counter (fec) 1.2.0, 2023-09-11

Searching...
Searching done.

Searching path          : e:\temp
Bypass permission error : true
Case sensitive          : true
Recursive               : true
Found file              : 3126
Found extension         : 79
Total size              : 5.063 GB
Elapsed time            : 69.1037ms

  No.  Extension        Count          Size

    1                     537    696.310 KB
    2  ._                  54       0 bytes
    3  .apk                 1     31.922 MB
    4  .at                 76    363.312 KB
    5  .bat                 7     15.746 KB
   ....
   69  .toml                4     711 bytes
   70  .txt               294      1.304 MB
   71  .TXT                28    482.062 KB
   72  .values              5      1.582 MB
   73  .vsix                1     56.140 KB
   74  .woff2               9   1004.168 KB
   75  .xaml                2     929 bytes
   76  .xls                 2    719.000 KB
   77  .xlsx                1     11.426 KB
   78  .xml               139     27.135 MB
   79  .zip                13      4.392 GB
```

### Search files extensions in case sensitive manner and sort result by size

```bash {.line-numbers}
$ fec -t -t -t -s e:/temp

Copyright (c) 1999-2023 Not a dream Co., Ltd.
file extension counter (fec) 1.2.0, 2023-09-11

Searching...
Searching done.

Searching path          : e:\temp
Bypass permission error : true
Case sensitive          : true
Recursive               : true
Found file              : 3126
Found extension         : 79
Total size              : 5.063 GB
Elapsed time            : 67.6833ms

  No.  Extension        Count          Size

    1  .zip                13      4.392 GB
    2  .rar                 1    351.305 MB
    3  .docx                2     76.881 MB
    4  .dll               120     46.956 MB
    5  .pdf                 6     37.749 MB
   ....
   69  .settings            3     699 bytes
   70  .properties         12     542 bytes
   71  .kotlin_module       4     498 bytes
   72  .code-workspace      1     410 bytes
   73  .gitattributes       1     160 bytes
   74  .MF                  4     144 bytes
   75  .java                1     144 bytes
   76  .probe               2      16 bytes
   77  .name                1      14 bytes
   78  ._                  54       0 bytes
   79  .cargo-lock          3       0 bytes
```
