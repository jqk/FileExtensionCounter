# FileExtensionCounter

FileExtensionCounter is a command line program used to search all file extensions under a given path, and give the number of files and total file size for each extension.

Only tested on Windows but should also run on Linux and Mac operating systems.

## Usage

```bash {.line-numbers}
$ fec

Copyright (c) 1999-2023 Not a dream Co., Ltd.
file extension counter (fec) 1.0.0, 2023-07-03

Usage:
  fec [command] <path/to/counter/extensions>

Command:
  -f, --false, or omit command: Count extensions in specified path, extension is case insensitive.
  -t, --true                  : Count extensions in specified path, extension is case sensitive.

  otherwise: show this help.
```

## Example

### Search files extensions in case insensitive manner

```bash {.line-numbers}
$ fec e:\temp

Copyright (c) 1999-2023 Not a dream Co., Ltd.
file extension counter (fec) 1.0.0, 2023-07-03

Searching path : e:\temp
Case sensitive : false
Found file     : 3527
Found extension: 83

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

Files with uppercase "TXT" are treated differently from files with lowercase "txt".

```bash {.line-numbers}
$ fec -t e:\temp

Copyright (c) 1999-2023 Not a dream Co., Ltd.
file extension counter (fec) 1.0.0, 2023-07-03

Searching path : e:\temp
Case sensitive : true
Found file     : 3527
Found extension: 84

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
