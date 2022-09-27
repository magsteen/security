#include <iostream>
#include <iterator>
#include <regex>
#include <string>

std::string swap(std::string text) {
    text = std::regex_replace(text, std::regex("&"), "&amp;");
    text = std::regex_replace(text, std::regex("<"), "&lt;");
    text = std::regex_replace(text, std::regex(">"), "&gt;");
    return text;
}

int main() {
    std::string test_text = "Foo<&>Bar";
    std::cout << "Initial: " << test_text << std::endl; 
    std::cout << "  After: " << swap(test_text) << std::endl; 
}
