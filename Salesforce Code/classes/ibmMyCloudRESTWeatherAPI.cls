public class ibmMyCloudRESTWeatherAPI {
    public void getdata(){
        List<Disaster__c> newDisasters = [Select Id From Disaster__c];
    }
    
    public void send(){
 Http http1=new Http();
 HttpRequest req1=new HttpRequest();
 //Our Test Endpoint Yours Will Differ
        String url = 'https://api.weather.com/v3/alerts/headlines?geocode=33.037582,-117.252616&format=json&language=en-US&apiKey=5424e9662cbf4bc3a4e9662cbf4bc3fe';
 
 req1.setendpoint(url);
 req1.Setheader('String','String');
 req1.setmethod('GET'); //you can also SET method `GET` with Get there is no need of req1.setbody()
 HttpResponse res1;
 res1 = http1.send(req1);
 String str=res1.getbody();
    }
}