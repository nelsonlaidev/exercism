import 'dart:math';

class HighScores {
  final List<int> scores;

  HighScores(this.scores);

  int latest() {
    return scores[scores.length - 1];
  }

  int personalBest() {
    return scores.reduce(max);
  }

  List<int> personalTopThree() {
    List<int> sorted = scores.toList()..sort((a, b) => b.compareTo(a));

    return sorted.sublist(0, min(3, scores.length));
  }
}
