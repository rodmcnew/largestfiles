This is a CLI program that displays the largest files and directories on your file system. This is written in Golang.

## Usage

Usage example:

```bash
# Display large files and directories inside the /Applications directory
largestfiles /Applications
```

Output example:

```
Looking in /Applications
Found 14.2 GB of usage in 189,648 files and 16,196 directories. Scanning took 2.301s

- - - - - - - - - - Largest Directories - - - - - - - - - -
    Size  Files  Path
503.5 MB    264  /Applications/Microsoft Word.app/Contents/Resources/DFonts
503.5 MB    264  /Applications/Microsoft PowerPoint.app/Contents/Resources/DFonts
503.5 MB    264  /Applications/Microsoft Excel.app/Contents/Resources/DFonts
370.0 MB      6  /Applications/Microsoft Edge.app/Contents/Frameworks/Microsoft Edge Framework.framework/Versions/94.0.992.38
327.6 MB    205  /Applications/PhpStorm.app/Contents/lib
306.5 MB     99  /Applications/Microsoft OneNote.app/Contents/Resources/DFonts
297.0 MB     15  /Applications/Firefox.app/Contents/MacOS
278.8 MB     98  /Applications/Microsoft Outlook.app/Contents/Resources/DFonts
198.4 MB     55  /Applications/VirtualBox.app/Contents/MacOS
177.0 MB      1  /Applications/Spotify.app/Contents/Frameworks/Chromium Embedded Framework.framework
171.7 MB      1  /Applications/Google Chrome.app/Contents/Frameworks/Google Chrome Framework.framework/Versions/94.0.4606.61
171.6 MB      1  /Applications/Google Chrome.app/Contents/Frameworks/Google Chrome Framework.framework/Versions/94.0.4606.71
162.8 MB     68  /Applications/krita.app/Contents/Frameworks
146.6 MB     65  /Applications/PhpStorm.app/Contents/jbr/Contents/Home/lib
140.6 MB      1  /Applications/Slack.app/Contents/Frameworks/Electron Framework.framework/Versions/A
131.4 MB      1  /Applications/Microsoft Teams.app/Contents/Frameworks/Electron Framework.framework/Versions/A
129.7 MB      1  /Applications/Visual Studio Code.app/Contents/Frameworks/Electron Framework.framework/Versions/A
127.5 MB    148  /Applications/krita.app/Contents/PlugIns
109.0 MB      1  /Applications/Microsoft Excel.app/Contents/MacOS
107.3 MB      2  /Applications/Microsoft Word.app/Contents/SharedSupport/Proofing Tools/FinnishGrammar.proofingtool/Contents/SharedSupport/FinnishGrammar.lexicon/Contents/Resources

- - - - - - - - - - - Largest Files - - - - - - - - - - - -
    Size  Path
358.3 MB  /Applications/Microsoft Edge.app/Contents/Frameworks/Microsoft Edge Framework.framework/Versions/94.0.992.38/Microsoft Edge Framework
276.3 MB  /Applications/Firefox.app/Contents/MacOS/XUL
177.0 MB  /Applications/Spotify.app/Contents/Frameworks/Chromium Embedded Framework.framework/Chromium Embedded Framework
171.7 MB  /Applications/Google Chrome.app/Contents/Frameworks/Google Chrome Framework.framework/Versions/94.0.4606.61/Google Chrome Framework
171.6 MB  /Applications/Google Chrome.app/Contents/Frameworks/Google Chrome Framework.framework/Versions/94.0.4606.71/Google Chrome Framework
140.6 MB  /Applications/Slack.app/Contents/Frameworks/Electron Framework.framework/Versions/A/Electron Framework
131.4 MB  /Applications/Microsoft Teams.app/Contents/Frameworks/Electron Framework.framework/Versions/A/Electron Framework
129.7 MB  /Applications/Visual Studio Code.app/Contents/Frameworks/Electron Framework.framework/Versions/A/Electron Framework
109.0 MB  /Applications/Microsoft Excel.app/Contents/MacOS/Microsoft Excel
107.3 MB  /Applications/Microsoft Outlook.app/Contents/SharedSupport/Proofing Tools/FinnishGrammar.proofingtool/Contents/SharedSupport/FinnishGrammar.lexicon/Contents/Resources/DataFile.lex
107.3 MB  /Applications/Microsoft Word.app/Contents/SharedSupport/Proofing Tools/FinnishGrammar.proofingtool/Contents/SharedSupport/FinnishGrammar.lexicon/Contents/Resources/DataFile.lex
 96.3 MB  /Applications/PhpStorm.app/Contents/lib/platform-impl.jar
 81.3 MB  /Applications/Microsoft Outlook.app/Contents/Frameworks/WordMail.framework/Versions/A/WordMail
 81.3 MB  /Applications/Microsoft Word.app/Contents/MacOS/Microsoft Word
 81.3 MB  /Applications/Microsoft Word.app/Contents/SharedSupport/Open XML for Excel.app/Contents/MacOS/Open XML for Excel
 81.3 MB  /Applications/Microsoft PowerPoint.app/Contents/SharedSupport/Open XML for Excel.app/Contents/MacOS/Open XML for Excel
 73.1 MB  /Applications/Microsoft Word.app/Contents/Frameworks/mso99.framework/Versions/A/mso99
 73.1 MB  /Applications/Microsoft Excel.app/Contents/Frameworks/mso99.framework/Versions/A/mso99
 73.1 MB  /Applications/Microsoft Outlook.app/Contents/Frameworks/mso99.framework/Versions/A/mso99
 73.1 MB  /Applications/Microsoft PowerPoint.app/Contents/Frameworks/mso99.framework/Versions/A/mso99
```

Usage examples with extra options:

```bash
# Show 50 results
largestfiles -c=50 /Applications

# Ignore file system errors such as "permission denied" and keep going
largestfiles -i /Applications
```

## Installation

```bash
go install github.com/rodmcnew/largestfiles/cmd/largestfiles@latest
```
