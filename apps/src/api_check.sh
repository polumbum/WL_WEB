#!/bin/bash

BGreen='\033[1;32m' 
Purple='\033[0;35m'       # Purple
NC='\033[0m' # No Color


echo -e "${BGreen}\nSPORTSMEN\n${NC}"

# Login sportsman
echo -e "${BGreen}curl -X POST -d '{"email":"ivan@mail.ru","password":"123"}' http://localhost:8000/users/login${NC}"

response=$(curl -X POST -d '{"email":"ivan@mail.ru","password":"123"}' http://localhost:8000/users/login)

token=$(echo "$response" | jq -r '.token')


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $token" http://localhost:8000/sportsmen${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $token" http://localhost:8000/sportsmen


echo -e "${BGreen}GET:${Purple} curl http://localhost:8000/sportsmen${NC}"

curl -w "%{http_code}\n" http://localhost:8000/sportsmen


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/?fullname=Давид+Оле${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/?fullname=Давид+Оле


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/?fullname=Олегович+Да${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/?fullname=Олегович+Да


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/?fullname=олегович+Да${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/?fullname=олегович+Да


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/?page=1&batch=2${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $token" "http://localhost:8000/sportsmen/?page=1&batch=2"


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/?page=1&batch=2&sort=name.desc${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $token" "http://localhost:8000/sportsmen/?page=1&batch=2&sort=name.desc"


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/?sort=name.desc${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $token" "http://localhost:8000/sportsmen/?sort=name.desc"


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115

# Login secretary
echo -e "${BGreen}curl -X POST -d '{"email":"sec@mail.ru","password":"123"}' http://localhost:8000/users/login${NC}"

response=$(curl -X POST -d '{"email":"sec@mail.ru","password":"123"}' http://localhost:8000/users/login)

tokenSec=$(echo "$response" | jq -r '.token')


echo -e "${BGreen}PATCH:${Purple} curl -X PATCH  curl -H "Authorization: Bearer $tokenSec" -d '{"id":"f8f26e6d-3e36-416a-9ff0-253876fc1115","sports_category":"МСМК"}' http://localhost:8000/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115${NC}"

curl -w "%{http_code}\n" -X PATCH -H "Authorization: Bearer $tokenSec" -d '{"id":"f8f26e6d-3e36-416a-9ff0-253876fc1115","sports_category":"МСМК"}' http://localhost:8000/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115/results${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $token" http://localhost:8000/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115/results


# echo -e "${BGreen}DELETE:${Purple} curl -X DELETE http://localhost:8000/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115${NC}"

# curl -X DELETE http://localhost:8000/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115


echo -e "${BGreen}POST:${Purple} curl -X POST -H "Authorization: Bearer $tokenSec" -d '{"c_id":"f8f26e6d-3e36-416a-9ff0-253876fc1111"}' http://localhost:8000/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115/coach${NC}"

curl -w "%{http_code}\n" -X POST -H "Authorization: Bearer $tokenSec" -d '{"c_id":"f8f26e6d-3e36-416a-9ff0-253876fc1111"}' http://localhost:8000/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115/coach


echo -e "${BGreen}\nCOACHES\n${NC}"

#Login coach
echo -e "${BGreen}curl -X POST -d '{"email":"alex@mail.ru","password":"123"}' http://localhost:8000/users/login${NC}"

response=$(curl -X POST -d '{"email":"alex@mail.ru","password":"123"}' http://localhost:8000/users/login)

tokenCoach=$(echo "$response" | jq -r '.token')


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $tokenCoach"  http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111/sportsmen/results${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $tokenCoach"  http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111/sportsmen/results


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $tokenCoach" http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111/sportsmen${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $tokenCoach" http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111/sportsmen


echo -e "${BGreen}DELETE:${Purple} curl -X DELETE -H "Authorization: Bearer $tokenSec" http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115${NC}"

curl -w "%{http_code}\n" -X DELETE -H "Authorization: Bearer $tokenSec" http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111/sportsmen/f8f26e6d-3e36-416a-9ff0-253876fc1115


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $tokenCoach" http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111/sportsmen${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $tokenCoach" http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111/sportsmen


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $tokenCoach" http://localhost:8000/coaches${NC}"

curl  -w "%{http_code}\n" -H "Authorization: Bearer $tokenCoach" http://localhost:8000/coaches


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $tokenCoach" http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $tokenCoach" http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111


# echo -e "${BGreen}DELETE:${Purple} curl -X DELETE http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111${NC}"

# curl -X DELETE http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111


# echo -e "${BGreen}GET:${Purple} curl http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111${NC}"

# curl http://localhost:8000/coaches/f8f26e6d-3e36-416a-9ff0-253876fc1111


echo -e "${BGreen}\nCOMPETITIONS\n${NC}"

#Login comp org
echo -e "${BGreen}curl -X POST -d '{"email":"comp@mail.ru","password":"123"}' http://localhost:8000/users/login${NC}"

response=$(curl -X POST -d '{"email":"comp@mail.ru","password":"123"}' http://localhost:8000/users/login)

tokenCompOrg=$(echo "$response" | jq -r '.token')

echo -e "${BGreen}GET:${Purple} curl http://localhost:8000/competitions/a8f26e6d-3e36-416a-9ff0-253876fc1115/results${NC}"

curl -w "%{http_code}\n" http://localhost:8000/competitions/a8f26e6d-3e36-416a-9ff0-253876fc1115/results


echo -e "${BGreen}GET:${Purple} curl http://localhost:8000/competitions${NC}"

curl -w "%{http_code}\n" http://localhost:8000/competitions


echo -e "${BGreen}POST:${Purple} curl -X POST -H "Authorization: Bearer $tokenCompOrg" -d '{
    "name":"Кубок Москвы",
    "city":"Москва",
    "address":"Московская 1",
    "beg_date":"2024-10-24T13:55:14.614Z",
    "end_date":"2024-10-24T13:55:14.614Z",
    "age":"Юниоры, юниорки",
    "min_sports_category":"КМС",
    "antidoping": true,
    "org_id":"a8f26e6d-3e36-416a-9ff0-253876fc1000"
    }' http://localhost:8000/competitions${NC}"

curl -w "%{http_code}\n" -X POST -H "Authorization: Bearer $tokenCompOrg" -d '{
    "name":"Кубок Москвы",
    "city":"Москва",
    "address":"Московская 1",
    "beg_date":"2024-10-24T13:55:14.614Z",
    "end_date":"2024-10-24T13:55:14.614Z",
    "age":"Юниоры, юниорки",
    "min_sports_category":"КМС",
    "antidoping": true,
    "org_id":"a8f26e6d-3e36-416a-9ff0-253876fc1000"
    }' http://localhost:8000/competitions


echo -e "${BGreen}POST:${Purple} curl -X POST -H "Authorization: Bearer $token" -d '{
  "sm_id": "f8f26e6d-3e36-416a-9ff0-253876fc1115",
  "weight_category": 109,
  "start_snatch": 100,
  "start_clean_and_jerk": 140
}' http://localhost:8000/competitions/a8f26e6d-3e36-416a-9ff0-253876fc1115/sportsman${NC}"

curl -w "%{http_code}\n" -X POST -H "Authorization: Bearer $token" -d '{
  "sm_id": "f8f26e6d-3e36-416a-9ff0-253876fc1115",
  "weight_category": 109,
  "start_snatch": 100,
  "start_clean_and_jerk": 140
}' http://localhost:8000/competitions/a8f26e6d-3e36-416a-9ff0-253876fc1115/sportsman


echo -e "${BGreen}DELETE:${Purple} curl -X DELETE -H "Authorization: Bearer $tokenCompOrg" http://localhost:8000/competitions/a8f26e6d-3e36-416a-9ff0-253876fc1115${NC}"

curl -w "%{http_code}\n" -X DELETE -H "Authorization: Bearer $tokenCompOrg" http://localhost:8000/competitions/a8f26e6d-3e36-416a-9ff0-253876fc1115


echo -e "${BGreen}GET:${Purple} curl http://localhost:8000/competitions${NC}"

curl -w "%{http_code}\n" http://localhost:8000/competitions


echo -e "${BGreen}\nTCAMPS\n${NC}"

echo -e "${BGreen}curl -X POST -d '{"email":"tcamp@mail.ru","password":"123"}' http://localhost:8000/users/login${NC}"

response=$(curl -X POST -d '{"email":"tcamp@mail.ru","password":"123"}' http://localhost:8000/users/login)

tokenTCampOrg=$(echo "$response" | jq -r '.token')

echo -e "${BGreen}GET:${Purple} curl http://localhost:8000/tcamps${NC}"

curl -w "%{http_code}\n" http://localhost:8000/tcamps


echo -e "${BGreen}POST:${Purple} curl -X POST -H "Authorization: Bearer $tokenTCampOrg" -d '{
    "city":"Москва",
    "address":"Московская 1",
    "beg_date":"2024-10-24T13:55:14.614Z",
    "end_date":"2024-10-24T13:55:14.614Z",
    "org_id":"a8f26e6d-3e36-416a-9ff0-253876fc1000"
    }' http://localhost:8000/tcamps${NC}"

curl -w "%{http_code}\n" -X POST -H "Authorization: Bearer $tokenTCampOrg" -d '{
    "city":"Москва",
    "address":"Московская 1",
    "beg_date":"2024-10-24T13:55:14.614Z",
    "end_date":"2024-10-24T13:55:14.614Z",
    "org_id":"a8f26e6d-3e36-416a-9ff0-253876fc1000"
    }' http://localhost:8000/tcamps


echo -e "${BGreen}POST:${Purple} curl -X POST -H "Authorization: Bearer $token" -d '{
  "sm_id": "f8f26e6d-3e36-416a-9ff0-253876fc1115"
}' http://localhost:8000/tcamps/aaa26e6d-3e36-416a-9ff0-253876fc1115/sportsman${NC}"

curl -w "%{http_code}\n" -X POST -H "Authorization: Bearer $token" -d '{
  "sm_id": "f8f26e6d-3e36-416a-9ff0-253876fc1115"
}' http://localhost:8000/tcamps/aaa26e6d-3e36-416a-9ff0-253876fc1115/sportsman


echo -e "${BGreen}DELETE:${Purple} curl -X DELETE -H "Authorization: Bearer $tokenTCampOrg" http://localhost:8000/tcamps/aaa26e6d-3e36-416a-9ff0-253876fc1115${NC}"

curl -w "%{http_code}\n" -X DELETE -H "Authorization: Bearer $tokenTCampOrg" http://localhost:8000/tcamps/aaa26e6d-3e36-416a-9ff0-253876fc1115


echo -e "${BGreen}GET:${Purple} curl http://localhost:8000/tcamps${NC}"

curl -w "%{http_code}\n" http://localhost:8000/tcamps


echo -e "${BGreen}\nUSERS\n${NC}"

echo -e "${BGreen}curl -X POST -d '{"email":"tcamp@mail.ru","password":"123"}' http://localhost:8000/users/login${NC}"

response=$(curl -X POST -d '{"email":"tcamp@mail.ru","password":"123"}' http://localhost:8000/users/login)

tokenTCampOrg=$(echo "$response" | jq -r '.token')



echo -e "${BGreen}POST:${Purple} curl -X POST -d '
{
  "email": "user@example.com",
  "password": "12345",
  "role": "secretary"
}' http://localhost:8000/users/signup${NC}"

curl -w "%{http_code}\n" -X POST -d '
{
  "email": "user@example.com",
  "password": "12345",
  "role": "secretary"
}' http://localhost:8000/users/signup


echo -e "${BGreen}POST:${Purple} curl -X POST -d '
{
  "email": "user@example.com",
  "password": "12345"
}' http://localhost:8000/users/login${NC}"

curl -w "%{http_code}\n" -X POST -d '
{
  "email": "user@example.com",
  "password": "12345"
}' http://localhost:8000/users/login


echo -e "${BGreen}GET:${Purple} curl http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbbbb${NC}"

curl -w "%{http_code}\n" http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbbbb


echo -e "${BGreen}PUT:${Purple} curl -X PUT -d '
{"email":"comp@mail.ru",
"password":"321"
}' http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbbbb${NC}"

curl -w "%{http_code}\n" -X PUT -d '
{"email":"comp@mail.ru",
"password":"321"
}' http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbbbb


echo -e "${BGreen}DELETE:${Purple} curl -X DELETE http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbbbb${NC}"

curl -w "%{http_code}\n" -X DELETE http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbbbb


echo -e "${BGreen}GET:${Purple} curl http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbbbb${NC}"

curl -w "%{http_code}\n" http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbbbb


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $tokenTCampOrg" http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbddd/tcamps${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $tokenTCampOrg" http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbddd/tcamps


echo -e "${BGreen}GET:${Purple} curl -H "Authorization: Bearer $tokenCompOrg" http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbccc/competitions${NC}"

curl -w "%{http_code}\n" -H "Authorization: Bearer $tokenCompOrg" http://localhost:8000/users/f8f26e6d-3e36-416a-9ff0-253876fcbccc/competitions

