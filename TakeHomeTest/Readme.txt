The URL verifier application is implemented using GO Language for the server side, HTML and JQuery for the client side.

Details for the Development Tools:

GO version: go1.10.3.windows-amd64.msi
OS: Windows 7 64 Bit
JQuery version: 3.1.1

The Main files to be looked upon 

url_check.go Located at (TakeHomeTest\src\url_verify)
url_check.html Located at (TakeHomeTest\src\url_verify)

URL to execute
http://localhost:9090/add_url


Unit Test Case Files
url_check_test.go Located at (TakeHomeTest\test\url_verify_test)

Each folders url_verify and url_verify_test contains a batch file exec.bat.
Running the batch files will run the application


Steps ro run the application

1. Download and Install go1.10.3.windows-amd64.msi
2. Copy the Directory TakeHomeTest in Desktop
3. Set the %PATH% in the in the environment variable for the   Directory containing the source
4. Go to following folder
   TakeHomeTest\src\url_verify\
5. Run exec.bat from there and keep it running.
6. Open the browser and enter the following URL
   https://localhost:9090/add_url

7. Enter URL one by one and click on check button.

8. A table with URL and statuses will be displayed.

9. If any URL is deleted it wont be visible any more.

10. For running the test cases please go to the following folder

    TakeHomeTest\test\url_verify_test\
    Execute the exec.bat file
    Results displayed in the command prompt



References followed:
https://tour.golang.org/welcome/
https://www.tutorialspoint.com/go/