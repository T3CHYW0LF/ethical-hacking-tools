`urls-get-http-status.sh`

Reads a list of URLs from a file and output the http response code.

**Setup:**
- Save the file to your local computer
- chmod u+x urls-get-http-status.sh

**Example Usage:**
`./urls-get-http-status $(pwd)"/urls" | tee -a urls-http-status
