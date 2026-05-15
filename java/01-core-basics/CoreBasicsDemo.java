public class CoreBasicsDemo {

    public static void main(String[] args) {
        compareEquality();
        integerCache();
        passByValue();
        stringBuilderVsString();
        overloadAndOverride();
    }

    private static void compareEquality() {
        System.out.println("=== == 和 equals ===");
        String a = new String("java");
        String b = new String("java");
        System.out.println("a == b -> " + (a == b));
        System.out.println("a.equals(b) -> " + a.equals(b));

        User u1 = new User(1, "Tom");
        User u2 = new User(1, "Tom");
        System.out.println("u1.equals(u2) -> " + u1.equals(u2));
        System.out.println("u1.hashCode() == u2.hashCode() -> " + (u1.hashCode() == u2.hashCode()));
        System.out.println();
    }

    private static void integerCache() {
        System.out.println("=== Integer 缓存与自动装箱 ===");
        Integer x = 127;
        Integer y = 127;
        Integer m = 128;
        Integer n = 128;
        System.out.println("127: x == y -> " + (x == y));
        System.out.println("128: m == n -> " + (m == n));
        System.out.println("128: m.equals(n) -> " + m.equals(n));
        System.out.println();
    }

    private static void passByValue() {
        System.out.println("=== Java 只有值传递 ===");
        int number = 10;
        changeNumber(number);
        System.out.println("changeNumber 后 number -> " + number);

        User user = new User(2, "Alice");
        changeUserName(user);
        System.out.println("changeUserName 后 user.name -> " + user.name);

        reassignUser(user);
        System.out.println("reassignUser 后 user.name -> " + user.name);
        System.out.println();
    }

    private static void changeNumber(int number) {
        number = 99;
    }

    private static void changeUserName(User user) {
        user.name = "Bob";
    }

    private static void reassignUser(User user) {
        user = new User(3, "Carol");
    }

    private static void stringBuilderVsString() {
        System.out.println("=== String 和 StringBuilder ===");
        String text = "a";
        text += "b";
        text += "c";
        System.out.println("String 拼接结果 -> " + text);

        StringBuilder builder = new StringBuilder("a");
        builder.append("b");
        builder.append("c");
        System.out.println("StringBuilder 拼接结果 -> " + builder);
        System.out.println();
    }

    private static void overloadAndOverride() {
        System.out.println("=== 重载与重写 ===");
        Printer printer = new SmartPrinter();
        printer.print("hello");
        printer.print(123);
        System.out.println();
    }

    static class Printer {
        void print(String text) {
            System.out.println("Printer prints text: " + text);
        }

        void print(int number) {
            System.out.println("Printer prints number: " + number);
        }
    }

    static class SmartPrinter extends Printer {
        @Override
        void print(String text) {
            System.out.println("SmartPrinter overrides text print: " + text);
        }
    }

    static class User {
        int id;
        String name;

        User(int id, String name) {
            this.id = id;
            this.name = name;
        }

        @Override
        public boolean equals(Object obj) {
            if (this == obj) {
                return true;
            }
            if (!(obj instanceof User other)) {
                return false;
            }
            return id == other.id && name.equals(other.name);
        }

        @Override
        public int hashCode() {
            int result = Integer.hashCode(id);
            result = 31 * result + name.hashCode();
            return result;
        }
    }
}
