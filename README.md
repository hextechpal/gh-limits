# gh-limits

For all the following experiments the flow remains the same

CreateBranch -> CreateFile -> CreatePr

| Details                               | Description                                        | Start               | End                 |
|---------------------------------------|----------------------------------------------------|---------------------|---------------------|
| Execute Flow with 0 break in between  | Secondary limit reaches after 10 requests          |                     |                     |
| Execute Flow with 5s break in between | All the requests successful                        | 2023/09/20 17:40:02 | 2023/09/20 17:54:27 |
| Execute Flow with 5s break in between | Failed on 22nd iteration, secondary limit exceeded | 2023/09/20 18:13:28 | 2023/09/20 18:20:01 |
| Execute Flow with 5s break in between | Show file differences that haven't been staged     | 2023/09/20 17:40:02 | 2023/09/20 17:54:27 |
| Execute Flow with 5s break in between | Show file differences that haven't been staged     | 2023/09/20 17:40:02 | 2023/09/20 17:54:27 |
