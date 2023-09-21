# gh-limits

For all the following experiments the flow remains the same

* Flow1 : CreateBranch -> CreateFile -> CreatePr
* Flow2 : ListPr -> MergePR

| Details                                          | Description                                                                                                                                            | Log file | Start               | End                 |
|--------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|----------|---------------------|---------------------|
| Execute Flow1 with 0 break in between            | Secondary limit reaches after 10 requests                                                                                                              |          |                     |                     |
| Execute Flow1 with 5s break in between (50 req)  | All the requests successful                                                                                                                            | exp3.log | 2023/09/20 17:40:02 | 2023/09/20 17:54:27 |
| Execute Flow1 with 5s break in between (100 req) | Failed on 75th iteration, secondary limit exceeded                                                                                                     | exp5.log | 2023/09/21 17:11:05 | 2023/09/21 17:33:00 |
| Execute Flow2 with 5s break in between (67 req)  | All 67 Requests successful, but tried manually creating aPR of GH it failed Pull request creation failed. Validation failed: was submitted too quickly | exp6.log | 2023/09/21 17:37:02 | 2023/09/21 17:44:29 |


### Conclusions

* The secondary limits only apply to create/merge pull request and not to thing like creating a branch, file etc
* 75 req per hour is the rate limit
* Seems like the secondary rate limits are per API too