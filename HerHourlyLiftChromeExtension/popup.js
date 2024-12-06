const API_URL = "http://localhost:8080/quotes/random"; // Your backend API URL

// Fetch random quote from backend
async function fetchQuote() {
  try {
    const response = await fetch(API_URL);
    const data = await response.json();
    return data.quote || "No quote available.";
  } catch (error) {
    console.error("Error fetching quote:", error);
    return "Error fetching quote.";
  }
}

// Show new quote and start progress bar
async function showNewQuote() {
  startProgressBar();
  const quote = await fetchQuote();
  document.getElementById('quote').textContent = quote;
}

// Progress bar animation
function startProgressBar() {
  const progress = document.querySelector('.progress');
  progress.style.width = '0%';
  setTimeout(() => {
    progress.style.width = '100%';
  }, 100); // Start animation after a short delay
}

// Toggle theme
document.getElementById('theme-toggle').addEventListener('change', () => {
  document.body.classList.toggle('dark-mode');
});

// Load quote on popup open
showNewQuote();
