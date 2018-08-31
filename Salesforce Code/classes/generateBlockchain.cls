public class generateBlockchain {

    public Boolean showContent { get; set; }

    public generateBlockchain() {
        showContent = true;
    }
 
    public PageReference toggleContent() {
        if(showContent){
            showContent = false;
        }
        else{
            showContent = true;
        }
        return null;
    }
    public void donationCallout(){
     HttpRequest req = new HttpRequest();
     req.setEndpoint('localhost:8888');
     req.setMethod('GET');
//Insert Blockchain Code Endpoint Method Here
  
     Http http = new Http();
     HTTPResponse res = http.send(req);
     System.debug(res.getBody());
   }
}