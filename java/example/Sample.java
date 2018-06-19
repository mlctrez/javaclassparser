package example;

public class Sample implements SampleInterface {

  private int sampleField;

  public Sample() {}

  public String sampleMethod(String arg) {
    java.util.ArrayList al = new java.util.ArrayList();
    System.out.println(al);
    return arg;
  }
}
