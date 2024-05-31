### StudentPlannerAPI
#### Giriş
Merhaba, bu API, bana gönderilen bir case için tarafımca yazılmıştır. Öğrencilerin planlarını yönetmelerine yardımcı olmak amacıyla tasarlanmıştır.
**İsteğe bağlı seçenecekler de dahil olmak üzere bütün beklentiler karşılanmıştır.**
#### Kullanılan Teknolojiler
* Go: API geliştirme için ana programlama dili.
* Gorm: Go için ORM (Object-Relational Mapping) kütüphanesi, veri tabanı işlemleri için kullanıldı.
* Echo: Go ile yazılmış bir HTTP framework, API endpoint'lerinin yönetimi için kullanıldı.
* MySQL: Veri tabanı yönetim sistemi olarak kullanıldı.

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