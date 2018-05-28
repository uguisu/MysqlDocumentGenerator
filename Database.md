
## GenderCode
_Gender code management_

| Column Name | Type | Length | Nullable | Column Comment |
| --------------- | --------------- | --------------- | --------------- | --------------- |
| `GENDER_CODE` | int |  | NO | Code |
| GENDER_NAME | varchar | 45 | NO | Name |
| GENDER_ORDER | tinyint |  | NO | Display order |

## User
_User basic information_

| Column Name | Type | Length | Nullable | Column Comment |
| --------------- | --------------- | --------------- | --------------- | --------------- |
| `USER_ID` | varchar | 6 | NO | User ID=CC-xxx |
| AGE | tinyint |  | YES | age |
| SEX | tinyint |  | YES | 0=Male; 1=Female; 2=Other; 3=unknow |
| USER_TYPE | tinyint |  | NO | 1=system; 2=wechat; 3=qq; 4=zhihu; 5=weibo; 6=linkedin |
