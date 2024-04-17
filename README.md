API GATEWAY ✅
USER SERVICE ✅
POST SERVICE ✅
COMMENT SERVICE ✅

user ni id bo'yicha get qilinganda yoki user list get qilinganda hamma postlar, postlarni commentlari va har bitta commentni owner'i haqida data qaytadi. ✅

- agar berilgan linklardan video topilmasa, videos zip tashlaganman, hamma video'lar bor nomlari bilan.

- inMemory, redis tool'lari bilan test qilish mumkin.
- user testing ✅
- post testing ✅
- comment testing ✅

Docker ✅
- dockerfile in all microservices ✅

Swagger ✅
- mock data to postgres in migrate up ✅

Mock microservice testing(mock file ichida, microservice lar mocklangan va har bir method uchun test yozilgan, bunda servicelarga connect qilinmedi): 
- user ✅
- post ✅
- comment ✅

Mongo database in microservices:
- user ✅ 
- post ✅ 
- comment ✅ 

Mongo database testing in microservices(mongo uchun test yozilgan har bir serviceda):
- user ✅
- post ✅
- comment ✅

**EXTRAs to exam tasks:**

Real microservice testing(bunda real microservicelarni test qiladi, microservice-testing file ichida, bu test ni run qilish uchun microservice lar ishlab turgan holda bo'lishi kerak): 
- user  ✅
- post ✅
- comment ✅
- (bu testing, microservice lardan rostan ham response kelyaptimi, tog'ri ishlayaptimi tekshirish maqsadda).

Message broker ✅ 
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

  

