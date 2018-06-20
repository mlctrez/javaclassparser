package example;

public class Sample implements SampleInterface {

  private int sampleField;

  public Sample() {}

  public String sampleMethod(String arg) {
    java.util.ArrayList al = new java.util.ArrayList();
    System.out.println(al);
    return arg;
  }

  public void sampleComplexMethod() {
    String foo = "bar";
    if (foo == null) {
      System.out.println("foo is null");
    }

    int baaz = 0;
    if (baaz > 10) {
      System.out.println("baaz > 10");
    }
  }
}
