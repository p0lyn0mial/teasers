package mapdecoder;

import java.util.Map;
import java.util.HashMap;
import java.util.regex.Pattern;

import java.util.Arrays;

public class MyMapDecoder implements MapDecoder {

      // Implement a method that decodes a String to a corresponding Map. The String is composed of key-value pairs. Key is separated from value with an 'equals' character ('='). Pairs are separated with an 'ampersand' character ('&'). Example: key=value&key=value&....
     // Empty keys and values are allowed, but the equals sign must be present (e.g. "=value", "key=").
     // If the key is empty, empty String should be put to the map as key.
     // If the value is empty, empty String should be put to the map as value.
     // If the given String is empty, an empty Map should be returned.
     // If the given String is null, null should be returned.
     // If the given String has an invalid format, an IllegalArgumentException should be thrown.
     // Your implementation must implement the MapDecoder interface.
    public Map<String, String> decode(String s) {
      if (s.isEmpty()) {
	      return new HashMap<String, String>();
      }
      if (s == null) {
	      return null;
      }
      // TODO: compile pattern
      if (!Pattern.matches("^[\\w=]*.*[\\w=]$", s)) {
            throw new IllegalArgumentException("Illegal string passed in:" + s);
      }


      Map<String, String> result = new HashMap<String, String>(); 
      String[] encodedKVPairs = s.split("&");

      for(String encodedKVPair : encodedKVPairs){
	 encodedKVPair = encodedKVPair.trim();
	 // TODO: compile pattern
	 if (!Pattern.matches("^\\w*+=\\w*$",encodedKVPair)) {
            throw new IllegalArgumentException("Illegal string passed in:" + s);
	 }

         String[] keyValuePair = encodedKVPair.split("=");
         if (keyValuePair.length == 0) { throw new IllegalArgumentException("Illegal string passed in: " + s); }
         if (keyValuePair.length > 2) { throw new IllegalArgumentException("Illegal string passed in: " + s); }

	 String key = keyValuePair[0];
	 String value = "";
	 if (keyValuePair.length > 1) {
		 value = keyValuePair[1];
	 }

	 //TODO: what about duplicates ?
	 result.put(key, value);
         
      }

      // an empty result means we were not able to find any valid input
      if (result.isEmpty()) {
            throw new IllegalArgumentException("Illegal string passed in: " + s);
      }
      return result;
    }

    public static void main(String[] args) {
      // scenario 1
      {
	Map<String, String> expectedMap = new HashMap<String, String>();
	expectedMap.put("one", "1");
	expectedMap.put("two","2");
        String toDecode = "one=1&two=2";
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	if (!decodedMap.equals(expectedMap)) {
          System.out.println("scenario 1 failed");
	  System.exit(-1);
	}
      }
      // scenario 2 
      {
	Map<String, String> expectedMap = new HashMap<String, String>();
	expectedMap.put("one", "1");
        String toDecode = "one=1";
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	if (!decodedMap.equals(expectedMap)) {
          System.out.println("scenario 2 failed");
	  System.exit(-1);
	}
        //for(Map.Entry entry:decodedMap.entrySet()){  
          //System.out.println(entry.getKey()+" "+entry.getValue());  
	//}
      }
      
      // scenario 3
      {
        String toDecode = "one=1=two=2";
	boolean exceptionOcurred = false; 
	try {
	  Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	} catch (IllegalArgumentException exp) {
          exceptionOcurred = true;
	}
	if (!exceptionOcurred) {
          System.out.println("scenario 3 failed");
	  System.exit(-1);
	}
      }
      // scenario 4
      {
        String toDecode = "&&";
	boolean exceptionOcurred = false; 
	try {
	  Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	} catch (IllegalArgumentException exp) {
          exceptionOcurred = true;
	}
	if (!exceptionOcurred) {
          System.out.println("scenario 4 failed");
	  System.exit(-1);
	}
      }
      // scenario 5
      {
        String toDecode = "=&";
	boolean exceptionOcurred = false; 
	try {
	  Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	} catch (IllegalArgumentException exp) {
          exceptionOcurred = true;
	}
	if (!exceptionOcurred) {
          System.out.println("scenario 5 failed");
	  System.exit(-1);
	}
      }
      // scenario 6
      {
	Map<String, String> expectedMap = new HashMap<String, String>();
	expectedMap.put("", "value");
        String toDecode = "=value";
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	if (!decodedMap.equals(expectedMap)) {
          System.out.println("scenario 6 failed");
	  System.exit(-1);
	}
      }
      // scenario 7
      {
        System.out.println("starting scenario 7 ");
	Map<String, String> expectedMap = new HashMap<String, String>();
	expectedMap.put("key", "");
        String toDecode = "key=";
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	if (!decodedMap.equals(expectedMap)) {
          System.out.println("scenario 7 failed");
	  System.exit(-1);
	}
      }
      // scenario 8
      {
        System.out.println("starting scenario 8 ");
	Map<String, String> expectedMap = new HashMap<String, String>();
	expectedMap.put("one", "1");
	expectedMap.put("two", "2");
	expectedMap.put("three", "3");
	expectedMap.put("four", "4");
        String toDecode = "one=1&two=2&three=3&four=4";
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	if (!decodedMap.equals(expectedMap)) {
          System.out.println("scenario 8 failed");
	  System.exit(-1);
	}
      }
      // scenario 9
      {
        System.out.println("starting scenario 9 ");
	boolean exceptionOcurred = false; 
        String toDecode = "one=1&two=2&&three=3&four=4";
	try {
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	} catch(IllegalArgumentException exp) {
          exceptionOcurred = true;
	}
	if (!exceptionOcurred) {
          System.out.println("scenario 9 failed");
	  System.exit(-1);
	}
      }
      // scenario 10
      {
        System.out.println("starting scenario 10 ");
	boolean exceptionOcurred = false; 
        String toDecode = "one=1&two=2=&three=3&four=4";
	try {
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	} catch(IllegalArgumentException exp) {
          exceptionOcurred = true;
	}
	if (!exceptionOcurred) {
          System.out.println("scenario 10 failed");
	  System.exit(-1);
	}
      }
      // scenario 11
      {
        System.out.println("starting scenario 11 ");
	boolean exceptionOcurred = false; 
        String toDecode = "one=1&two=2=&&three=3&four=4";
	try {
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	} catch(IllegalArgumentException exp) {
          exceptionOcurred = true;
	}
	if (!exceptionOcurred) {
          System.out.println("scenario 11 failed");
	  System.exit(-1);
	}
      }
      // scenario 12
      {
        System.out.println("starting scenario 12 ");
	boolean exceptionOcurred = false; 
        String toDecode = "one=1&two=2=&=&three=3&four=4";
	try {
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	} catch(IllegalArgumentException exp) {
          exceptionOcurred = true;
	}
	if (!exceptionOcurred) {
          System.out.println("scenario 12 failed");
	  System.exit(-1);
	}
      }
      // scenario 13
      {
        System.out.println("starting scenario 13 ");
	boolean exceptionOcurred = false; 
        String toDecode = "one=1&two=2&";
	try {
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	} catch(IllegalArgumentException exp) {
          exceptionOcurred = true;
	}
	if (!exceptionOcurred) {
          System.out.println("scenario 13 failed");
	  System.exit(-1);
	}
      }
      // scenario 14
      {
        System.out.println("starting scenario 14 ");
	boolean exceptionOcurred = false; 
        String toDecode = "one=1&=";
	try {
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	} catch(IllegalArgumentException exp) {
          exceptionOcurred = true;
	}
	if (!exceptionOcurred) {
          System.out.println("scenario 14 failed");
	  System.exit(-1);
	}
      }
      // scenario 15
      {
        System.out.println("starting scenario 15 ");
	Map<String, String> expectedMap = new HashMap<String, String>();
	expectedMap.put("one", "1");
	expectedMap.put("", "value");
        String toDecode = "one=1&=value";
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	if (!decodedMap.equals(expectedMap)) {
          System.out.println("scenario 15 failed");
	  System.exit(-1);
	}
      }
      // scenario 16
      {
        System.out.println("starting scenario 16 ");
	Map<String, String> expectedMap = new HashMap<String, String>();
	expectedMap.put("one", "1");
	expectedMap.put("key", "");
        String toDecode = "one=1&key=";
	Map<String, String> decodedMap = new MyMapDecoder().decode(toDecode);
	if (!decodedMap.equals(expectedMap)) {
          System.out.println("scenario 16 failed");
	  System.exit(-1);
	}
      }
    }
}

