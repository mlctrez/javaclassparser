package example;

public class OtherSample {

  private SampleInterface iface;

  public String viaInterface(String arg) {
    return iface.sampleMethod(arg);
  }

  public String viaDirec(String arg) {
    return new Sample().sampleMethod(arg);
  }

  public void accessPublicConstant() {
    System.out.println(Sample.PUB_CONST);
  }


}
