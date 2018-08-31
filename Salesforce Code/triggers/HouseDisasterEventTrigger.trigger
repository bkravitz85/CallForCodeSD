/**
* ─────────────────────────────────────────────────────────────────────────────────────────────────┐
* Code Review: IBM Apex App Connect Platform Event
* Class Name: appConnect_DisasterEvent
* Description: Executes upsert of disaster data driven from current API
* ──────────────────────────────────────────────────────────────────────────────────────────────────
* @secondaryDeveloper         TEST     test@test@ibm.com
* @Tester     Tester   email
* @ReviewedBy   Brandon Kravitz    brandon.kravitz@bluewolfgroup.com
* @userStory        US-0001
* @created         2017-11-13
* @modified        2017-11-13
* @systemLayer    Test
* @thisCodeReviewTestClass            appConnect_DisasterEvent
* @see            AppConnect Documentation
* ──────────────────────────────────────────────────────────────────────────────────────────────────
* @changes
* vX.X             brandon.kravitz@ibm.com
* YYYY-MM-DD      Initial Commit
* ─────────────────────────────────────────────────────────────────────────────────────────────────┘
*/		

// Trigger for listening to HouseDisaster__e events.
trigger HouseDisasterEventTrigger on HouseDisaster__e(after insert) {    
    // List to hold all Disasters to be created.
    List<Disaster__c> disasters = new List<Disaster__c>();
    
    // Documentation ID       
    // Iterate through each notification.
    for (HouseDisaster__e event : Trigger.New) {
        JSONParser parser = JSON.createParser(event.Disaster__c);
        while (parser.nextToken() != null) {
        // Start at the array of invoices.
        if (parser.getCurrentToken() == JSONToken.START_ARRAY) {
            while (parser.nextToken() != null) {
                // Advance to the start object marker to
                //  find next invoice statement object.
                if (parser.getCurrentToken() == JSONToken.START_OBJECT) {
                    // Read entire invoice object, including its array of line items.
                    DisasterClass d = (DisasterClass)parser.readValueAs(DisasterClass.class);
                    system.debug('long:: ' + d.longitude + 'lat:: ' + d.latitude );
                    system.debug('Size of list items: ' + d.x_id);
                    // For debugging purposes, serialize again to verify what was parsed.
                    String s = JSON.serialize(d);
                    system.debug('Serialized Disaster: ' + s);

                    // Skip the child start array and start object markers.
                    parser.skipChildren();
                    
                    Disaster__c disaster = new Disaster__c();
 if (d.x_id != null) disaster.Document_ID__c = d.x_id;
 if (d.satellite != null) disaster.Satellite__c = d.satellite ;
 if (d.bright_ti4 != null) disaster.Bright_ti4__c = Decimal.valueOf(d.bright_ti4);
 if (d.bright_ti5 != null) disaster.bright_ti5__c = Decimal.valueOf(d.bright_ti5);
 if (d.scan != null) disaster.scan__c = Decimal.valueOf(d.scan); 
 if (d.track != null) disaster.track__c = Decimal.valueOf(d.track); 
 if (d.confidence != null) disaster.confidence__c = d.confidence;
 if (d.longitude != null) disaster.longitude__c = Decimal.valueOf(d.longitude); 
 if (d.latitude != null) disaster.latitude__c = Decimal.valueOf(d.latitude); 
 if (d.latitude != null && d.longitude != null) disaster.Geolocation__longitude__s = Decimal.valueOf(d.longitude);
  if (d.latitude != null && d.longitude != null) disaster.Geolocation__latitude__s = Decimal.valueOf(d.latitude);
                   
                    disasters.add(disaster);
                }
            }
        }
      //  Map<String, Object> m = (Map<String, Object>)JSON.deserializeUntyped(event.Disaster__c);

   }
    
upsert disasters; 
    // Insert all Disasters corresponding to events received.
}
}