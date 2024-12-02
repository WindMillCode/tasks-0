bool isEmailValid(String email) {
  // Define a regular expression pattern for a basic email validation
  final RegExp emailRegex = RegExp(
    r'^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$',
  );

  return emailRegex.hasMatch(email);
}
