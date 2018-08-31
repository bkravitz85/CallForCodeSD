trigger DisasterEventTrigger on Disaster__c (before insert, before update) {
    List<Account> housedisasters  =  new List<Account>();

    Disaster__c d = Trigger.New[0];   
    
    housedisasters = [SELECT Id, Name, ShippingAddress, DISTANCE(ShippingAddress, GEOLOCATION(:d.Latitude__c, :d.Longitude__c), 'mi') FROM Account WHERE DISTANCE(ShippingAddress, GEOLOCATION(:d.Latitude__c, :d.Longitude__c), 'mi') < 100 and ShippingStreet <> NULL];
    
    
    System.debug('List of House Disasters:::: ' + housedisasters);
}