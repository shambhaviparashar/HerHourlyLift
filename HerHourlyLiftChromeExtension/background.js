const API_URL = "http://localhost:8080/quotes/random"; // Backend API URL

// Fetch quote from backend
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

// Show notification with quote
async function showNotification() {
  const quote = await fetchQuote();
  chrome.notifications.create({
    type: "basic",
    iconUrl: "icon128.png",
    title: "HerHourlyLift",
    message: quote
  });
}

// Set up hourly alarm
chrome.runtime.onInstalled.addListener(() => {
  chrome.alarms.create("hourlyAlarm", { periodInMinutes: 60 });
});

// Listen for the alarm and show notification
chrome.alarms.onAlarm.addListener((alarm) => {
  if (alarm.name === "hourlyAlarm") {
    showNotification();
  }
});
