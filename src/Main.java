import java.util.ArrayList;
import java.util.List;

class Main {

  public static void main(String[] args) {
    List<Integer> numbers = new ArrayList<Integer>();

    numbers.add(1);
    numbers.add(7);

    System.out.printf("List: %s \n", numbers.get(1));
  }
}
