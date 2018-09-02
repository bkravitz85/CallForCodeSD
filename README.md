# CallForCode Donate Your Unused Data to Firerelief Efforts

## Igniting Our Inspiration


![eap](https://media.giphy.com/media/vJzUgsALABTkk/giphy.gif)

On August 19th the planet celebrated the United Nation's Humanitarian Day to raise awareness of first responders who provide aid immediately after disasters.

> Call for Code aims to help improve the lives of those communities who are most threatened by natural disasters and in turn, will help ease and improve the important role that humanitarian agencies play.

Our project utilizes 5 myCloud Services and provides data donations to disaster first responders. Our system uses IBM Data Integration Services to send data to and from Salesforce, Machine Learning to find most at risk areas and Blockchain to record asset transfers. (Data)


![datasci](https://i1.wp.com/tellyourdatastory.com/wp-content/uploads/2018/08/datasci.png?w=980&ssl=1)
## Items We Used

| Item | Location | Description |
| --- | --- | --- |
| Salesforce | login.salesforce.com | Developer Edition of Salesforce with Analtyics Studio |
| IBM myCloud | https://console.bluemix.net/catalog/ | IBM's Catalog of Services |
| Weather Company API | http://developer.weather.com | Raw Data to Use for Competition CSV Data File and JSON Data from the API | APP Connect | https://console.bluemix.net/catalog/ | Allows to integrate IBM Data to Salesforce

## HIC MANEBIMUS OPTIME

In order to build an event driven disaster project, lets define a few metrics so that our datasets can be meaningful and perform to our scope:

| **Define Scope:** | **Predict where first responders will have their data throttled....** |
| --- | --- |
| **Define Performance:** | **Our target accurancy should be over < 70%** |
| **Context: ** | **Using IBM Weather Company Data along with Wigle.net API we can predict with 70% or greater accuracy which areas will have deficient bandwidth resources (Cell, Wifi, etc) .** |
| Solution Data:  | Using Machine Learning Workflows to process and transform Weather Company Data to create a prediction model, this model must predict which homes/devices will be effected by impending types of disasters. |

## Creating Object Model

Let's take our proposed data types and transform them into some metadata!

First we are going to define a few items:

| Item | Description |
| --- | --- |
| Household (Account) | Describes homes with Geolocation, Value, Income, Gender |
| Disaster__c | Disaster Types with Geolocation, Intensity, Time |
| HouseholdDisaster__e | Platform Event that can return values and transform multiple objects |

![pe](https://i0.wp.com/tellyourdatastory.com/wp-content/uploads/2018/08/pe.gif)

## Implementation

```java
/**
 * ────────────────────────────────────────────────────────────────────────────────────────────────
 * Code Review: IBM Apex App Connect Platform Event
 * Class Name: appConnect_DisasterEvent
 * Description: Executes upsert of disaster data driven from current API
 * ────────────────────────────────────────────────────────────────────────────────────────────────
 * @ReviewedBy   Brandon Kravitz    brandon.kravitz@bluewolfgroup.com
 * @thisCodeReviewTestClass            appConnect_DisasterEvent
 **/

// Trigger for listening to HouseDisaster__e events.
trigger HouseDisasterEventTrigger on HouseDisaster__e(after insert) {
 // List to hold all Disasters to be created.
 List < Disaster__c > disasters = new List < Disaster__c > ();

 // Documentation ID       
 // Iterate through each notification.
 for (HouseDisaster__e event: Trigger.New) {
  // Let's parse through our long text field and grab all disaster objects
  JSONParser parser = JSON.createParser(event.Disaster__c);
  while (parser.nextToken() != null) {
   if (parser.getCurrentToken() == JSONToken.START_ARRAY) {
    while (parser.nextToken() != null) {
     if (parser.getCurrentToken() == JSONToken.START_OBJECT) {
      //  System Debug Disasters
      DisasterClass d = (DisasterClass) parser.readValueAs(DisasterClass.class);
      system.debug('GEOLOCATION long:: ' + d.longitude + 'lat:: ' + d.latitude);
      system.debug('Size of list items: ' + d.x_id);
      String s = JSON.serialize(d);
      system.debug('Serialized Disaster: ' + s);

      // Skip the child start array and start object markers.
      parser.skipChildren();

      Disaster__c disaster = new Disaster__c();
      if (d.x_id != null) disaster.Document_ID__c = d.x_id;
      if (d.satellite != null) disaster.Satellite__c = d.satellite;
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
  }
  // Insert all Disasters corresponding to events received.    
  upsert disasters;
 }
}
```

## Integration
![disasterFields](https://i1.wp.com/tellyourdatastory.com/wp-content/uploads/2018/08/disasterFields.png?w=980&ssl=1)
Data retention and integration is the hallmark of our project. Leveraging several powerful platforms supercharges the ability for a call for code project to send and recieve data. To store the bulk of our unfiltered data, we utilized IBM Cloudant No SQL database to house our call for code project data. Our data utilized IBM's new APP Connect Service to send and recieve data from Cloudant to Salesforce, including IBM's Weather API, Wigle.net Devices and NASA's VISSA Fire data. 
App Connect makes it easy to create bulk data sends from NASA data, via a Callable Flow, Scheduled Send or API Extension.
We've plugged APP Connect into Salesforce Platform Events, which allows us to make alert notifications in real time, which in turn trigger geolocation distances, that will make it available to be analyzed.

![flowHD](https://i0.wp.com/tellyourdatastory.com/wp-content/uploads/2018/08/flowHD.gif?w=980&ssl=1)
##  Keep Going!!!
