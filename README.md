# Uygulama Geliştirme Süresi 
tüm geliştirme ve araştırma ile harcadığım zaman 6 saattir. Aynı zamanda çalıştığım için boş zamanlarımda yapabildim.

go öğrendikten sonra direk framework ile yazmaya başladım bu yuzden bazı şeyleri tekrar gözden geçirdiğim için uzun zaman aldı ama yeni şeyler keşfetmek de faydalı oldu.

# Ornek url Adresleri

## veri set etmek için 

http://localhost:9001/set?key=userID&value=968

http://localhost:9001/set?key=isLogin&value=true

http://localhost:9001/set?key=headerOption&value=exist

## veri get etmek için 

http://localhost:9001/get?key=userID

http://localhost:9001/get?key=isLogin

http://localhost:9001/get?key=headerOption

## tüm veriyi ekrana yazar

http://localhost:9001/print

# Unit test 
testleri bilerek yazmadım.


# Golang Ödevi - Test Case 1

Sizden bir web/api server uygulaması yazmanızı istiyoruz. Bu uygulama, hafızada çalışan bir depo olacak ve bilgileri key/value şeklinde saklayacak. 
Özetle sizden istediğimiz basit bir In-memory key/value store uygulaması. 
Uygulamanın iki adet end-point’i olacak; 

set : gelen değeri atamak için; localhost:9001/set? key=ball&value=top

get : atanan değeri okumak için; localhost:9001/get?key=ball 

Ek olarak, belli zaman aralıklarında da, hafızadaki key/value database’ini fiziksel olarak dosyaya;
örneğin /tmp/TIMESTAMP-db.txt gibi, yazan bir otomasyon ( go routine ?) istiyoruz.  
Tek kısıtımız, sadece go ’nun standart kütüphanesinin kullanılmasını istiyoruz,
   
    yani 3. parti paket kullanmadan uygulama geliştirilmeli. 
    
    Bunun haricinde istediğiniz gibi projeyi şekillendirebilirsiniz, 
    
    http handler method’undan tutun, aklınıza gelen her fikri uygulamakta serbestsiniz.


# Yararlandığım kaynaklar 

https://github.com/peterbourgon/diskv/blob/master/index.go

https://github.com/naqvijafar91/cuteDB

https://flexiple.com/key-value-javascript/

https://github.com/recoilme/slowpoke/blob/e21936cf1c6e0798e64551be3b2e3d6dc7681d5e/slowpoke.go#L779

https://www.practical-go-lessons.com/chap-32-templates# hestiacp-golang-templates
