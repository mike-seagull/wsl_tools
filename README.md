Tools for Windows Subsystem Linux (WSL)<br/>
<h4>Environment Variables</h4>

* UNIX_C = windows drive from wsl (eg. `/mnt/c`)
* UNIX_HOME = home in wsl (eg. `~` or `$HOME`)
* WINDOWS_C = primary windows drive in windows (eg. `C:\`)
* WINDOWS_HOME = home in windows (eg. `${WINDOWS_C}/Users/ThisGuyNeedsABeer`)
* SUBLIME_TEXT_UNIX = sublime text executable from wsl (eg. `${UNIX_C}/Program\ Files/Sublime\ Text\ 3/subl.exe`)
<h4>sublime_text</h4>
<h5>To open sublime text to the directory or file</h5>
<code>sublime_text $FILE_OR_DIRECTORY</code>
<h4>open</h4>
<p>opens file explorer to the directory or file</p>
<code>open $FILE_OR_DIRECTORY</code>
