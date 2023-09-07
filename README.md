# FileExtensionCounter

FileExtensionCounter is a command line program used to search all file extensions under a given path, and give the number of files and total file size for each extension.

Only tested on Windows but should also run on Linux and Mac operating systems.

FileExtensionCounter 是一个命令行程序，用于搜索给定路径下的所有文件扩展名，并给出每个扩展名的文件数量和总文件大小。该程序仅在 Windows 上测试过，但也应该可以在 Linux 和 Mac 操作系统上运行。

## Usage

```bash {.line-numbers}
$ fec

Copyright (c) 1999-2023 Not a dream Co., Ltd.
file extension counter (fec) 1.1.1, 2023-09-07

Usage:
  fec [command] <path/to/counting/extensions>
       counting extensions in specified path

Command:
  the first char of the command defines if the extension is case sensitive.
      't' is true, 'f' is false.
  the second one defines how to sort the result.
      'c' means sort by count.
      'n' means sort by extension.
      's' means sort by size.

  -fn: default command, can be omitted. case insensitive and sort the result by extension.
  -fc: case insensitive and sort the result by count.
  -fs: case insensitive and sort the result by size.
  -tn: case sensitive and sort the result by extension.
  -tc: case sensitive and sort the result by count.
  -ts: case sensitive and sort the result by size.

  otherwise: show this help.
```

## Example

### Search files extensions in case insensitive manner

```bash {.line-numbers}
$ fec e:\temp
Or
$ fec -fn e:\temp

Copyright (c) 1999-2023 Not a dream Co., Ltd.
file extension counter (fec) 1.1.1, 2023-09-07

Searching...
Searching done.

Searching path : e:\temp
Case sensitive : false
Found file     : 3527
Found extension: 83
Elapsed time   : 123.126ms

  No.  Extension        Count          Size

    1                     914    846.661 KB
    2  .0                   1      41 bytes
    3  ._                  54       0 bytes
    4  .apk                 1     31.922 MB
    5  .at                 76    363.312 KB
    6  .bat                 7     15.746 KB
    7  .bin               113     11.603 MB
    8  .cargo-lock          3       0 bytes
    9  .class             162    464.832 KB
   10  .code-workspace      1     410 bytes
   11  .config              5      3.169 KB
   12  .cs                 20     47.387 KB
   13  .csproj              3     18.392 KB
   14  .css                 7     61.475 KB
   15  .csv                 8    303.735 KB
   16  .d                  41     37.212 KB
   17  .dll               120     46.956 MB
   18  .docx                2     76.881 MB
   19  .emmx                2     27.021 KB
   20  .exe                11      3.368 MB
   21  .gitattributes       1     160 bytes
   22  .gitignore          13      2.400 KB
   23  .go                  3      9.240 KB
   24  .gradle              2     908 bytes
   25  .html                9     84.726 KB
   26  .iml                 4      1.139 KB
   27  .jar                 9    419.868 KB
   28  .java                1     144 bytes
   29  .jpg                 9      3.129 MB
   30  .js                 20     78.247 KB
   31  .json               55     26.799 KB
   32  .keystream          76    336.000 KB
   33  .kotlin_module       4     498 bytes
   34  .kt                 23     40.804 KB
   35  .kts                 5      7.348 KB
   36  .len               228      1.781 KB
   37  .lock               45      2.416 KB
   38  .log                15      2.290 MB
   39  .manifest            1      3.322 KB
   40  .md                 16     95.249 KB
   41  .mf                  4     144 bytes
   42  .mod                 1     209 bytes
   43  .name                1      14 bytes
   44  .nupkg              19     14.694 MB
   45  .o                 398      3.786 MB
   46  .p7s                19    332.941 KB
   47  .pb                  3      3.604 KB
   48  .pdb                11     18.190 MB
   49  .pdf                 6     37.749 MB
   50  .png                34    447.822 KB
   51  .probe               2      16 bytes
   52  .properties         12     542 bytes
   53  .py                 39    193.953 KB
   54  .pyc                31     80.668 KB
   55  .rar                 2    792.295 MB
   56  .resx                5     27.803 KB
   57  .rlib               10      6.997 MB
   58  .rmeta              30      5.639 MB
   59  .rs                  6      6.271 KB
   60  .s                   5      1.254 KB
   61  .sample             77    134.984 KB
   62  .settings            3     699 bytes
   63  .sln                 2      2.732 KB
   64  .sql                 3    752.736 KB
   65  .sum                 1      1.788 KB
   66  .suo                 2    239.500 KB
   67  .svg                 8     12.954 KB
   68  .tab                83    320.023 KB
   69  .tab_i              76      2.375 MB
   70  .tag                 5     919 bytes
   71  .tar                 6     32.793 MB
   72  .timestamp          36      1.688 KB
   73  .toml                4     711 bytes
   74  .txt               322      1.775 MB
   75  .values              5      1.582 MB
   76  .vsix                1     56.140 KB
   77  .woff2               9   1004.168 KB
   78  .xaml                2     929 bytes
   79  .xls                 2    719.000 KB
   80  .xlsx                1     11.426 KB
   81  .xml               139     27.135 MB
   82  .yml                 1      2.992 KB
   83  .zip                11      3.565 GB
```

### Search files extensions in case sensitive manner

Files with uppercase extension `.TXT` are treated differently from files with lowercase `.txt`.

```bash {.line-numbers}
$ fec -tn e:\temp

Copyright (c) 1999-2023 Not a dream Co., Ltd.
file extension counter (fec) 1.1.1, 2023-09-07

Searching...
Searching done.

Searching path : e:\temp
Case sensitive : true
Found file     : 3527
Found extension: 83
Elapsed time   : 123.126ms

  No.  Extension        Count          Size

    1                     914    846.661 KB
    2  .0                   1      41 bytes
    3  ._                  54       0 bytes
    4  .apk                 1     31.922 MB
    5  .at                 76    363.312 KB
    6  .bat                 7     15.746 KB
    7  .bin               113     11.603 MB
    8  .cargo-lock          3       0 bytes
    9  .class             162    464.832 KB
   10  .code-workspace      1     410 bytes
   11  .config              5      3.169 KB
   12  .cs                 20     47.387 KB
   13  .csproj              3     18.392 KB
   14  .css                 7     61.475 KB
   15  .csv                 8    303.735 KB
   16  .d                  41     37.212 KB
   17  .dll               120     46.956 MB
   18  .docx                2     76.881 MB
   19  .emmx                2     27.021 KB
   20  .exe                11      3.368 MB
   21  .gitattributes       1     160 bytes
   22  .gitignore          13      2.400 KB
   23  .go                  3      9.240 KB
   24  .gradle              2     908 bytes
   25  .html                9     84.726 KB
   26  .iml                 4      1.139 KB
   27  .jar                 9    419.868 KB
   28  .java                1     144 bytes
   29  .jpg                 9      3.129 MB
   30  .js                 20     78.247 KB
   31  .json               55     26.799 KB
   32  .keystream          76    336.000 KB
   33  .kotlin_module       4     498 bytes
   34  .kt                 23     40.804 KB
   35  .kts                 5      7.348 KB
   36  .len               228      1.781 KB
   37  .lock               45      2.416 KB
   38  .log                15      2.290 MB
   39  .manifest            1      3.322 KB
   40  .md                 16     95.249 KB
   41  .MF                  4     144 bytes
   42  .mod                 1     209 bytes
   43  .name                1      14 bytes
   44  .nupkg              19     14.694 MB
   45  .o                 398      3.786 MB
   46  .p7s                19    332.941 KB
   47  .pb                  3      3.604 KB
   48  .pdb                11     18.190 MB
   49  .pdf                 6     37.749 MB
   50  .png                34    447.822 KB
   51  .probe               2      16 bytes
   52  .properties         12     542 bytes
   53  .py                 39    193.953 KB
   54  .pyc                31     80.668 KB
   55  .rar                 2    792.295 MB
   56  .resx                5     27.803 KB
   57  .rlib               10      6.997 MB
   58  .rmeta              30      5.639 MB
   59  .rs                  6      6.271 KB
   60  .s                   5      1.254 KB
   61  .sample             77    134.984 KB
   62  .settings            3     699 bytes
   63  .sln                 2      2.732 KB
   64  .sql                 3    752.736 KB
   65  .sum                 1      1.788 KB
   66  .suo                 2    239.500 KB
   67  .svg                 8     12.954 KB
   68  .tab                83    320.023 KB
   69  .tab_i              76      2.375 MB
   70  .TAG                 5     919 bytes
   71  .tar                 6     32.793 MB
   72  .timestamp          36      1.688 KB
   73  .toml                4     711 bytes
   74  .txt               294      1.304 MB
   75  .TXT                28    482.062 KB
   76  .values              5      1.582 MB
   77  .vsix                1     56.140 KB
   78  .woff2               9   1004.168 KB
   79  .xaml                2     929 bytes
   80  .xls                 2    719.000 KB
   81  .xlsx                1     11.426 KB
   82  .xml               139     27.135 MB
   83  .yml                 1      2.992 KB
   84  .zip                11      3.565 GB
```

### Search files extensions in case sensitive manner and sort result by size

```bash {.line-numbers}
$ fec -ts e:\temp

Copyright (c) 1999-2023 Not a dream Co., Ltd.
file extension counter (fec) 1.1.1, 2023-09-07

Searching...
Searching done.

Searching path : e:\temp
Case sensitive : true
Found file     : 3527
Found extension: 84
Elapsed time   : 84.2557ms

  No.  Extension        Count          Size

    1  .zip                11      3.565 GB
    2  .rar                 2    792.295 MB
    3  .docx                2     76.881 MB
    4  .dll               120     46.956 MB
    5  .pdf                 6     37.749 MB
    6  .tar                 6     32.793 MB
    7  .apk                 1     31.922 MB
    8  .xml               139     27.135 MB
    9  .pdb                11     18.190 MB
   10  .nupkg              19     14.694 MB
   11  .bin               113     11.603 MB
   12  .rlib               10      6.997 MB
   13  .rmeta              30      5.639 MB
   14  .o                 398      3.786 MB
   15  .exe                11      3.368 MB
   16  .jpg                 9      3.129 MB
   17  .tab_i              76      2.375 MB
   18  .log                15      2.291 MB
   19  .values              5      1.582 MB
   20  .txt               294      1.304 MB
   21  .woff2               9   1004.168 KB
   22                     914    846.661 KB
   23  .sql                 3    752.736 KB
   24  .xls                 2    719.000 KB
   25  .TXT                28    482.062 KB
   26  .class             162    464.832 KB
   27  .png                34    447.822 KB
   28  .jar                 9    419.868 KB
   29  .at                 76    363.312 KB
   30  .keystream          76    336.000 KB
   31  .p7s                19    332.941 KB
   32  .tab                83    320.023 KB
   33  .csv                 8    303.735 KB
   34  .suo                 2    239.500 KB
   35  .py                 39    193.953 KB
   36  .sample             77    134.984 KB
   37  .md                 16     95.249 KB
   38  .html                9     84.726 KB
   39  .pyc                31     80.668 KB
   40  .js                 20     78.247 KB
   41  .css                 7     61.475 KB
   42  .vsix                1     56.140 KB
   43  .cs                 20     47.387 KB
   44  .kt                 23     40.804 KB
   45  .d                  41     37.212 KB
   46  .resx                5     27.803 KB
   47  .emmx                2     27.021 KB
   48  .json               55     26.799 KB
   49  .csproj              3     18.392 KB
   50  .bat                 7     15.746 KB
   51  .svg                 8     12.954 KB
   52  .xlsx                1     11.426 KB
   53  .go                  3      9.240 KB
   54  .kts                 5      7.348 KB
   55  .rs                  6      6.271 KB
   56  .pb                  3      3.604 KB
   57  .manifest            1      3.322 KB
   58  .config              5      3.169 KB
   59  .yml                 1      2.992 KB
   60  .sln                 2      2.732 KB
   61  .lock               45      2.416 KB
   62  .gitignore          13      2.400 KB
   63  .sum                 1      1.788 KB
   64  .len               228      1.781 KB
   65  .timestamp          36      1.688 KB
   66  .s                   5      1.254 KB
   67  .iml                 4      1.139 KB
   68  .xaml                2     929 bytes
   69  .TAG                 5     919 bytes
   70  .gradle              2     908 bytes
   71  .toml                4     711 bytes
   72  .settings            3     699 bytes
   73  .properties         12     542 bytes
   74  .kotlin_module       4     498 bytes
   75  .code-workspace      1     410 bytes
   76  .mod                 1     209 bytes
   77  .gitattributes       1     160 bytes
   78  .MF                  4     144 bytes
   79  .java                1     144 bytes
   80  .0                   1      41 bytes
   81  .probe               2      16 bytes
   82  .name                1      14 bytes
   83  ._                  54       0 bytes
   84  .cargo-lock          3       0 bytes
```
