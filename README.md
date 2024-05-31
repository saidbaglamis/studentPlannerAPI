### StudentPlannerAPI
#### Giriş
Merhaba. Bu API, bana gönderilen bir case için tarafımca yazılmıştır. Öğrencilerin planlarını yönetmelerine yardımcı olmak amacıyla tasarlanmıştır.
**İsteğe bağlı seçenekler de dahil olmak üzere case içindeki bütün beklentiler karşılanmıştır.**
#### Kullanılan Teknolojiler
* **Go**: API geliştirme için ana programlama dili.
* **Gorm**: Go için ORM (Object-Relational Mapping) kütüphanesi, veri tabanı işlemleri için kullanıldı.
* **Echo**: Go ile yazılmış bir HTTP framework, API endpoint'lerinin yönetimi için kullanıldı.
* **MySQL**: Veri tabanı yönetim sistemi olarak kullanıldı.

### API Hakkında
StudentPlannerAPI, öğrencilerin belirli tarih aralıklarında planlar oluşturmasını, bu planlar üzerinde değişiklik yapmasını ve planları yönetmesini sağlar. Ayrıca yeni öğrenciler eklenebilir ve mevcut öğrenciler silinebilir, öğrencilerle birlikte tüm planları da silinebilir.

### Özellikler
* Plan Oluşturma: Öğrenciler belirli tarih aralıklarında planlar oluşturabilirler.
* Zaman Çakışma Kontrolü: Oluşturulan planlar, zaman çakışmalarına karşı kontrol edilir.
* Plan Düzenleme: Mevcut planlar üzerinde değişiklik yapılabilir.
* Plan Silme: Oluşturulan planlar silinebilir.
* Öğrenci Yönetimi: Yeni öğrenciler eklenebilir ve mevcut öğrenciler silinebilir. 
* Bir öğrenci silindiğinde, ona ait tüm planlar da silinir.
* Listeleme: Seçilen öğrencinin planları haftalık veya aylık olarak listelenebilir.

### Proje dizini yapısı
```
studentPlannerAPI/
├── cmd/
│   └── main.go
├── internal/
│    ├── database/
│    │   └── database.go     
│    ├── handlers/
│    │   └── plan_handler.go
│    │   └── student_handler.go
│    ├── models/
│    │   └── plan.go
│    │   └── student.go
│    └── routes/
│        └── routes.go
│    
├── go.mod
└── go.sum

```

### Kullanım Hakkında
**internal/database/database.go dosyasındaki:**
```
dsn := "root:password@tcp(127.0.0.1:3306)/plannerDB?charset=utf8mb4&parseTime=True&loc=Local"
```
**"plannerDB" yazan yeri kendi makinenizdeki çalışmak istediğiniz database ismiyle değiştirip bağlantı kurmanız gerekiyor**

**İlk çalıştırmada eğer veritabanınız boş ise MySQL Workbench veya benzeri bir program ile tablolara manuel girdi yapabilir ya da Postman kullanarak POST request body ile girdi sağlayıp veritabanınızı doldurabilirsiniz.**
##### Sizin için örnek bir JSON bırakıyorum
```
{
    "Email": "ornek@posta.com",
    "Name": "Said Baglamis",
    "SchoolNumber": "1453",
    "Class": "12/A",
    "Plans": [
        {
            "StudentID": 1,
            "PlanInfo": "OOP lab ödevi",
            "PlanStatus": "iptal",
            "StartDate": "2024-06-01T16:00:00+03:00",
            "EndDate": "2024-06-01T17:00:00+03:00"

        }
    ]
}
```
#### Endpointler
```
        //student handlerlar
GET	    "/student"          //Bütün öğrencileri ve bilgilerini döndürür
GET	    "/student/:id"      //ID'si verilen öğrenci ve bilgilerini döndürür
POST        "/student"          //Öğrenci oluşturur
PUT	    "/student/:id"      //ID'si verilen öğrenci bilgilerini günceller
DELETE	    "/student/:id"      //ID'si verilen öğrenciyi ve bilgilerini siler

	//plan handlerlar
GET	    "/plan/:id"         //ID'si verilen plan bilgilerini döndürür
GET	    "/weekly/:id"       //ID'si verilen haftalık plan bilgilerini döndürür
GET	    "/monthly/:id"      //ID'si verilen aylık plan bilgilerini döndürür
POST	    "/plan"             //Manuel plan oluşturur
PUT	    "/plan/:id"         //ID'si verilen plan bilgilerini günceller
DELETE	    "/plan/:id"         //ID'si verilen planı siler
```
### Faydalı Kaynaklar

Bu API'yi yazarken karşılaştığım ve sizin de karşılaşabileceğiniz hata ve sorunları gidermede yararlı bulduğum içerik ve platformlar:
* https://stackoverflow.com/questions/10299148/mysql-error-1045-28000-access-denied-for-user-billlocalhost-using-passw
* https://stackoverflow.com/questions/21944936/error-1045-28000-access-denied-for-user-rootlocalhost-using-password-y
* https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-20-04
* https://dev.to/envitab/how-to-build-an-api-using-go-ffk
* https://dev.to/ayoubzulfiqar/go-the-ultimate-folder-structure-6gj
* https://www.programiz.com/golang/data-types
* https://pkg.go.dev/time