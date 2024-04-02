API GATEWAY ✅
USER CRUD ✅
POST CRUD ✅
COMMENT CRUD ✅

user ni id bo'yicha get qilinganda yoki user list get qilinganda hamma postlar, postlarni commentlari va har bitta commentni owner'i haqida data qaytadi. ✅

- agar berilgan linklardan video topilmasa, videos zip tashlaganman, hamma video'lar bor nomlari bilan.

api-testing: video -> https://pub-daa7217568964be2861c94926a89ccab.r2.dev/fe4acd04-095e-4d40-983e-23e046874578.mp4
- inMemory, redis tool'lari bilan test qilish mumkin.
- user testing ✅
- post testing ✅
- comment testing ✅

Docker ✅
- dockerfile in all microservices ✅

Swagger ✅
- mock data to postgres in migrate up ✅

Mock microservice testing(mock file ichida, microservice lar mocklangan va har bir method uchun test yozilgan, bunda servicelarga connect qilinmedi): video -> https://pub-daa7217568964be2861c94926a89ccab.r2.dev/d13f72bd-ca87-4ea3-833d-bd114b35ff21.mp4
- user ✅
- post ✅
- comment ✅

Mongo database in microservices:
- user ✅ video -> https://pub-daa7217568964be2861c94926a89ccab.r2.dev/a73e5088-e4d3-43e3-adff-29dd390a5355.mp4
- post ✅ video -> https://pub-daa7217568964be2861c94926a89ccab.r2.dev/f28b8df5-bd36-46a3-b71e-74a9c1402627.mp4
- comment ✅ video -> https://pub-daa7217568964be2861c94926a89ccab.r2.dev/26d2c824-1c7d-45dd-9633-4ae5b8419224.mp4

Mongo database testing in microservices(mongo uchun test yozilgan har bir serviceda):
- user ✅
- post ✅
- comment ✅

**EXTRAs to exam tasks:**

Real microservice testing(bunda real microservicelarni test qiladi, microservice-testing file ichida, bu test ni run qilish uchun microservice lar ishlab turgan holda bo'lishi kerak): video -> https://pub-daa7217568964be2861c94926a89ccab.r2.dev/7b5bccab-1ecb-411f-ab46-7a9509072b13.mp4 
- user  ✅
- post ✅
- comment ✅
- (bu testing, microservice lardan rostan ham response kelyaptimi, tog'ri ishlayaptimi tekshirish maqsadda).

Message broker ✅ video -> https://pub-daa7217568964be2861c94926a89ccab.r2.dev/a58e13d2-c9c9-4e93-9652-c29c7d85c6b1.mp4
- kafka qo'shilgan -> create user method kafka bilan ishlaydi, api gateway user data sini produce qiladi, user service bo'sa consume qilib database'ga qo'shadi.

Unit postgres test in microservices(har bir microservice storage ga test yozilgan) - bularni videolari videos zip ni ichida bor:
- user ✅
- post ✅
- comment ✅

Cassandra database(yangi database qo'shilgan) - bularni videolari videos zip ni ichida bor:
- user ✅
- post ✅
- user ✅
- har bir microservice da cassandra implementation qilingan

Cassandra database testing in microservices(cassandra uchun test yozilgan har bir serviceda) - bularni videolari videos zip ni ichida bor:
- user ✅
- post ✅
- comment ✅


Kafka testing ✅ - bularni videolari videos zip ni ichida bor
- kafka-testing folder da, kafka ga test yozilgan api da produce ga test yozilgan, user service da bo'sa consume ga test yozilgan

Redis testing ✅: - bularni videolari videos zip ni ichida bor
- redis testing file da, redis ga test yozilgan

Casbin(middleware) ✅ - bularni videolari videos zip ni ichida bor
- casbin qo'shilgan
- admin (create, get, delete methodlar qo'shilgan, admin management uchun)

  

