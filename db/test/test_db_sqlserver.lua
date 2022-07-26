local db = require("db")

print("123123")

local sqlserver, err = db.open("sqlserver", "Provider=SQLOLEDB;Data Source=192.168.200.66\\gy;Initial Catalog=yxhis;User ID=sa;Password=123qwe,.;", { shared = true })
if err ~= nil then
    print(err)
end

local result, err = sqlserver:query("select * from yxhis.dbo.tbzdbq")
if err ~= nil then
    print(err)
end

-- for _, row in pairs(result.rows) do
    
--     print("row="..row)
-- end

for k, v  in pairs(result.columns) do
    print("k="..k.." v="..v)
end

for k, v in pairs(result.rows) do
    print(k)
    print(v[2].. " " ..v[6])
end

updateName = "update yxhis.dbo.tbzdbq set cmc = '心血管内科+1' where ibm = 3"
local result, err = sqlserver:exec(updateName)
if err ~= nil then
    print(err)
end

print("=============================================")

print(result.last_insert_id)
print(result.rows_affected)