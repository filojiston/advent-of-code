import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;
import java.util.stream.Collectors;

public class Main {
    private static final int PARTICLE_COLLISION_LOOPS = 1000;

    public static void main(String[] args) {
        String input = "";
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException exception) {
            System.err.println("Error reading input file");
        }

        List<Particle> particles = parseInput(input);
        System.out.println("Part 1: " + part1(particles));
        System.out.println("Part 2: " + part2(particles));
    }

    private static int part1(List<Particle> particles) {
        int closestParticle = 0;
        int closestDistance = Integer.MAX_VALUE;
        for (int i = 0; i < particles.size(); i++) {
            Particle particle = particles.get(i);
            int distance = Math.abs(particle.acceleration.x) + Math.abs(particle.acceleration.y)
                    + Math.abs(particle.acceleration.z);

            if (distance < closestDistance) {
                closestParticle = i;
                closestDistance = distance;
            }
        }
        return closestParticle;
    }

    private static int part2(List<Particle> particles) {
        for (int i = 0; i < PARTICLE_COLLISION_LOOPS; i++) {
            particles.stream().forEach(particle -> particle.tick());
            List<Particle> collisions = particles.stream().filter(particle -> {
                return particles.stream().filter(otherParticle -> {
                    return !particle.equals(otherParticle)
                            && particle.position.equals(otherParticle.position);
                }).count() > 0;
            }).collect(Collectors.toList());
            particles.removeAll(collisions);
        }
        return particles.size();
    }


    private static List<Particle> parseInput(String input) {
        return input.lines().map(line -> {
            String[] parts = line.split(", ");
            String[] positionParts = parts[0].substring(3, parts[0].length() - 1).split(",");
            String[] velocityParts = parts[1].substring(3, parts[1].length() - 1).split(",");
            String[] accelerationParts = parts[2].substring(3, parts[2].length() - 1).split(",");

            Vector3D position = new Vector3D(Integer.parseInt(positionParts[0]),
                    Integer.parseInt(positionParts[1]), Integer.parseInt(positionParts[2]));
            Vector3D velocity = new Vector3D(Integer.parseInt(velocityParts[0]),
                    Integer.parseInt(velocityParts[1]), Integer.parseInt(velocityParts[2]));
            Vector3D acceleration = new Vector3D(Integer.parseInt(accelerationParts[0]),
                    Integer.parseInt(accelerationParts[1]), Integer.parseInt(accelerationParts[2]));

            return new Particle(position, velocity, acceleration);
        }).collect(Collectors.toList());
    }

    private static class Particle {
        Vector3D position;
        Vector3D velocity;
        Vector3D acceleration;

        public Particle(Vector3D position, Vector3D velocity, Vector3D acceleration) {
            this.position = position;
            this.velocity = velocity;
            this.acceleration = acceleration;
        }

        public void tick() {
            velocity.x += acceleration.x;
            velocity.y += acceleration.y;
            velocity.z += acceleration.z;

            position.x += velocity.x;
            position.y += velocity.y;
            position.z += velocity.z;
        }

        @Override
        public String toString() {
            return "Particle{" + "position=" + position + ", velocity=" + velocity
                    + ", acceleration=" + acceleration + '}';
        }
    }

    private static class Vector3D {
        int x, y, z;

        public Vector3D(int x, int y, int z) {
            this.x = x;
            this.y = y;
            this.z = z;
        }

        @Override
        public String toString() {
            return "<" + x + ", " + y + ", " + z + '>';
        }

        @Override
        public boolean equals(Object obj) {
            if (obj instanceof Vector3D) {
                Vector3D other = (Vector3D) obj;
                return x == other.x && y == other.y && z == other.z;
            }
            return false;
        }
    }
}
