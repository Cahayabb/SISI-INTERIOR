# Fix API Errors Task

## Status: In Progress

### Steps:
- [x] 1. Copy favicon.ico from frontend to backend uploads/ for serving
- [ ] 2. Start ML prediction server (python sisi-interior-system/machine_learning/predict_api.py)
- [ ] 3. Refactor sisi-interior-system/controllers/estimasi_controllers.go: Fix GetEstimasi() - add ML error handling/fallback, replace ioutil.ReadAll, handle resp errors
- [ ] 4. Update sisi-interior-system/main.go: Add /favicon.ico static serve
- [ ] 5. Restart Go server and test endpoints (no 404/500)
- [ ] 6. Update TODO.md as complete

Current step: 2/6 (retry)
