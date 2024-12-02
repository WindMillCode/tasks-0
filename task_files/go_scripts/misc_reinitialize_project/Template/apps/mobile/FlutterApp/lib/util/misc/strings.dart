class TextEditor {
  String text;
  List<String> lines;

  TextEditor({required this.text}) : lines = text.split('\n');

  void insert(int line, int char, String newText) {
    line = line.clamp(0, lines.length);

    if (line >= lines.length) {
      lines.addAll(List<String>.filled(line - lines.length + 1, ''));
    }

    String targetLine = lines[line];

    int safeCharIndex = char.clamp(0, targetLine.length);

    lines[line] = targetLine.substring(0, safeCharIndex) +
        newText +
        targetLine.substring(safeCharIndex);

    text = lines.join('\n');
  }

  void replace(
      int fromLine, int fromChar, int toLine, int toChar, String newText) {
    // Clamp the line indices to valid ranges
    fromLine = fromLine.clamp(0, lines.length - 1);
    toLine = toLine.clamp(0, lines.length - 1);

    // Clamp the character indices to valid ranges within the specified lines
    fromChar = fromChar.clamp(0, lines[fromLine].length);
    toChar = toChar.clamp(0, lines[toLine].length);

    // Calculate the actual start index in the full text
    var fullText = lines.join('\n');
    var startIndex = lines.take(fromLine).join('\n').length +
        (fromLine > 0 ? 1 : 0) +
        fromChar;
    var endIndex =
        lines.take(toLine).join('\n').length + (toLine > 0 ? 1 : 0) + toChar;

    // Replace the specified range with the new text
    fullText = fullText.substring(0, startIndex) +
        newText +
        fullText.substring(endIndex);

    // Update the lines by splitting the modified full text
    lines = fullText.split('\n');
    text = fullText;
  }

  Map<String, dynamic> getTextDetails() {
    return {
      'totalLines': lines.length,
      'lines': lines
          .asMap()
          .entries
          .map((entry) =>
              {'lineNumber': entry.key + 1, 'charCount': entry.value.length})
          .toList()
    };
  }
}
