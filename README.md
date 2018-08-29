# CallForCode Part I / IV - Building Awesome Disaster Machine Learning Applications for ~~Dummies~~ (Me && You)

## Telling Your Datastory

Anyone in tech knows that now more than ever companies are changing their business requirements and compliance structure on a more frequent basis. With Deep Learning Forward Propagating Networks and constant advances in autonomous services; integrating data realtime is a competitive advantage any skilled developer should have in their wheelhouse. Working with Event Driven Architecture and on-demand data doesn't have to be overly complex, we will look at two classic platforms Salesforce and IBM Bluemix; apply Event Driven architecture (Platform Events) through several methodoligies and use cases.

## IBM Call for Code

https://media.giphy.com/media/vJzUgsALABTkk/giphy.gif

![eap](https://media.giphy.com/media/vJzUgsALABTkk/giphy.gif)![eap](https://media.giphy.com/media/vJzUgsALABTkk/giphy.gif)

On August 19th the planet celebrated the United Nation's Humanitarian Day to raise awareness of first responders who provide aid immediately after disasters.

> Call for Code aims to help improve the lives of those communities who are most threatened by natural disasters and in turn, will help ease and improve the important role that humanitarian agencies play.

Projects must utilize 5 myCloud Services and provide a need to disaster first responders. Sample projects included utilizing aerial drones that could provide topography insights into which areas need aid, uncovering deep learning image processing to treat and identify victims, while using the provided IBM "Weather Company API" to fuel disaster data and realness to the competition.

## Things You Will Need

| Item | Location | Description |
| --- | --- | --- |
| Salesforce | login.salesforce.com | Developer Edition of Salesforce with Analtyics Studio |
| IBM myCloud | https://console.bluemix.net/catalog/ | IBM's Catalog of Services |
| Weather Company API | http://developer.weather.com | Raw Data to Use for Competition CSV Data File and JSON Data from the API |

## HIC MANEBIMUS OPTIME

In order to build an event driven disaster project, lets define a few metrics so that our datasets can be meaningful and perform to our scope:

| **Define Scope:** | **Predict if a household disaster will occur...** |
| --- | --- |
| **Define Performance:** | **Our target accurancy should be over < 70%** |
| **Context: ** | **Using IBM Weather Company Data we can predict with 70% or greater accuracy which homes will be effected by disasters.** |
| Solution Data:  | Using Machine Learning Workflows to process and transform Weather Company Data to create a prediction model, this model must predict which homes will be effected by impending types of disasters. |

As long as we keep our scope within reasonable limits we should be able to create two distinct objects, Disasters and Households.

## So you wanna be a Salesforce Enterprise Architect?

![eap](eap.gif)![eap](file:///Users/kr4v/Documents/eap.gif)

Let's take our proposed data types and transform them into some metadata!

First we are going to define a few items:

| Item | Description |
| --- | --- |
| Household (Account) | Describes homes with Geolocation, Value, Income, Gender |
| Disaster__c | Disaster Types with Geolocation, Intensity, Time |
| HouseholdDisaster__e | Platform Event that can return values and transform multiple objects |

![pe](/Users/kr4v/Documents/pe.gif)![pe](file:///Users/kr4v/Documents/pe.gif)

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

## Code

##  Wrap Up
