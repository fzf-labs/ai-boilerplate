---
name: prototype-design
description: Prototype design tool for generating high-fidelity UI/UX prototypes from PRD documents. Trigger when users need to create interactive prototypes based on product requirements, validate UI/UX designs, or quickly generate prototype pages for requirement confirmation. Supports both mobile and PC applications with multiple design systems (WeChat Work, iOS Native, Material Design, Ant Design Mobile).
---

# Prototype Design

## Overview

Transform Product Requirements Documents (PRDs) into high-fidelity, interactive HTML/CSS/JavaScript prototypes. Generate production-ready prototype code that accurately reflects requirements, supports multiple design systems, and enables rapid design validation and stakeholder feedback.

## Core Capabilities

1. **PRD-to-Prototype Conversion**: Automatically parse PRD documents and generate corresponding UI prototypes
2. **Multi-Platform Support**: Create prototypes for mobile (iOS/Android) and PC (web) applications
3. **Design System Integration**: Support WeChat Work, iOS Native, Material Design, and Ant Design Mobile
4. **Interactive Prototypes**: Generate clickable, navigable prototypes with realistic interactions
5. **Responsive Design**: Ensure prototypes work across different screen sizes and devices
6. **Rapid Iteration**: Enable quick design changes based on feedback

## Workflow

### Step 1: Analyze PRD Document

**Read and understand the PRD:**

1. **Locate PRD file**: Check `docs/` directory for PRD documents
   ```bash
   # Common PRD locations
   docs/{date}-{feature-name}/prd.md
   docs/*/prd.md
   ```

2. **Extract key information**:
   - **User personas**: Who will use this feature?
   - **User stories**: What actions do users need to perform?
   - **Functional requirements**: What features must be included?
   - **UI/UX preferences**: Any specific design guidelines?
   - **Technical constraints**: Platform, device, performance requirements
   - **Success metrics**: How will we measure success?

3. **Identify prototype scope**:
   - Determine which pages/screens to prototype
   - Identify primary user flows to demonstrate
   - Note any specific interactions or animations
   - Understand MVP vs. future enhancements

**Ask clarifying questions if needed:**
- "Which user flow should I prioritize for the prototype?"
- "Do you have any design references or mockups?"
- "What level of interactivity do you need? (static, clickable, fully interactive)"
- "Should I focus on mobile or desktop first?"

### Step 2: Select Design System

**Choose appropriate design system based on:**

**Platform & Context:**
- **Mobile iOS app** → iOS Native Style
- **Mobile Android app** → Material Design Style
- **Chinese enterprise app** → WeChat Work Style
- **Enterprise mobile forms** → Ant Design Mobile Style
- **Cross-platform web** → Material Design or custom

**User Demographics:**
- **Chinese business users** → WeChat Work
- **Western consumers** → iOS Native or Material Design
- **Enterprise users** → WeChat Work or Ant Design Mobile

**Application Type:**
- **Consumer social/content app** → iOS Native or Material Design
- **Enterprise B2B tool** → WeChat Work or Ant Design Mobile
- **E-commerce** → Material Design or custom
- **Dashboard/analytics** → Ant Design Mobile or Material Design

**If not specified in PRD**, recommend a design system and explain the rationale.

**Design system characteristics:**

| Design System | Primary Color | Typography | Key Features |
|--------------|---------------|------------|--------------|
| WeChat Work | Tech Blue #3478F6 | PingFang SC | Simple, professional, enterprise-focused |
| iOS Native | System Blue #007AFF | SF Pro | Minimalist, spacious, Apple ecosystem |
| Material Design | Purple #6200EE | Roboto | Bold, elevation, ripple effects |
| Ant Design Mobile | Blue #108EE9 | SF UI Display | Efficient, form-heavy, business apps |

### Step 3: Design Page Structure

**Map PRD requirements to UI components:**

**Common page patterns:**

**1. Dashboard/Home Page:**
- Top navigation bar (title, search, menu)
- Quick access grid (icon buttons for key features)
- Data summary cards (metrics, KPIs)
- Feature list (navigable items)
- Bottom tab bar (mobile) or sidebar (desktop)

**2. List/Browse Page:**
- Search and filter bar
- List items with thumbnails, titles, metadata
- Pagination or infinite scroll
- Empty state for no results
- Pull-to-refresh (mobile)

**3. Detail Page:**
- Header with back button and actions
- Content sections (text, images, data)
- Related items or recommendations
- Action buttons (edit, delete, share)

**4. Form/Input Page:**
- Form fields grouped by section
- Input validation and error messages
- Submit and cancel buttons
- Progress indicator for multi-step forms

**5. Settings/Profile Page:**
- User info section
- Grouped settings list
- Toggle switches, dropdowns, navigation items
- Logout or destructive actions at bottom

**Create page hierarchy:**
```
Application
├─ Home/Dashboard
│  ├─ Quick Actions
│  ├─ Summary Cards
│  └─ Feature List
├─ Feature A
│  ├─ List View
│  └─ Detail View
├─ Feature B
│  ├─ Form View
│  └─ Confirmation View
└─ Settings
   ├─ Profile
   └─ Preferences
```

### Step 4: Generate Prototype Code

**Create HTML structure with embedded CSS and JavaScript:**

**File organization:**
- **Single-page prototype**: `prototype.html` (all pages in one file with tab/navigation switching)
- **Multi-page prototype**: `index.html`, `feature-a.html`, `feature-b.html`, etc.
- **Component-based**: Separate HTML files with shared `styles.css` and `script.js`

**Code structure template:**

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>[Feature Name] Prototype</title>

    <!-- Tailwind CSS CDN (optional) -->
    <script src="https://cdn.tailwindcss.com"></script>

    <!-- Custom Tailwind Config -->
    <script>
        tailwind.config = {
            theme: {
                extend: {
                    colors: {
                        'primary': '#3478F6',
                        'secondary': '#576B95',
                        // ... design system colors
                    }
                }
            }
        }
    </script>

    <!-- FontAwesome for icons (optional) -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">

    <style>
        /* Custom styles */
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB', sans-serif;
            background-color: #F5F5F5;
        }

        /* Add design system-specific styles */
    </style>
</head>
<body>
    <!-- Prototype content -->

    <script>
        // Interactive behaviors
        // Page navigation
        // Form validation
        // State management
    </script>
</body>
</html>
```

**Implementation guidelines:**

**1. Use semantic HTML:**
```html
<header>...</header>
<nav>...</nav>
<main>...</main>
<section>...</section>
<footer>...</footer>
```

**2. Apply design system styles consistently:**
- Use exact color codes from design system
- Follow spacing and sizing conventions
- Implement proper typography hierarchy
- Add appropriate shadows and borders

**3. Add realistic content:**
- Use actual feature names and descriptions from PRD
- Include realistic data (numbers, names, dates)
- Show multiple states (empty, loading, error, success)
- Provide sample data that reflects real use cases

**4. Implement interactions:**
- Button click handlers
- Form input validation
- Page/tab navigation
- Modal dialogs
- Dropdown menus
- Accordion sections
- Toast notifications

**5. Ensure responsiveness:**
```css
/* Mobile-first approach */
.container {
    width: 100%;
    max-width: 375px; /* iPhone width */
    margin: 0 auto;
}

/* Tablet */
@media (min-width: 768px) {
    .container {
        max-width: 768px;
    }
}

/* Desktop */
@media (min-width: 1200px) {
    .container {
        max-width: 1200px;
    }
}
```

### Step 5: Add Interactivity

**Implement key interactions:**

**1. Navigation:**
```javascript
// Tab navigation
function switchTab(tabId) {
    // Hide all tabs
    document.querySelectorAll('.tab-content').forEach(tab => {
        tab.classList.add('hidden');
    });

    // Show selected tab
    document.getElementById(tabId).classList.remove('hidden');

    // Update active state
    document.querySelectorAll('.tab-button').forEach(btn => {
        btn.classList.remove('active');
    });
    event.target.classList.add('active');
}
```

**2. Form validation:**
```javascript
function validateForm(formId) {
    const form = document.getElementById(formId);
    const inputs = form.querySelectorAll('input[required]');

    let isValid = true;
    inputs.forEach(input => {
        if (!input.value.trim()) {
            input.classList.add('error');
            isValid = false;
        } else {
            input.classList.remove('error');
        }
    });

    return isValid;
}
```

**3. Modal dialogs:**
```javascript
function showModal(modalId) {
    document.getElementById(modalId).classList.remove('hidden');
}

function hideModal(modalId) {
    document.getElementById(modalId).classList.add('hidden');
}
```

**4. Toast notifications:**
```javascript
function showToast(message, type = 'success') {
    const toast = document.createElement('div');
    toast.className = `toast toast-${type}`;
    toast.textContent = message;
    document.body.appendChild(toast);

    setTimeout(() => {
        toast.remove();
    }, 3000);
}
```

**5. Loading states:**
```javascript
function showLoading() {
    document.getElementById('loading').classList.remove('hidden');
}

function hideLoading() {
    document.getElementById('loading').classList.add('hidden');
}

// Simulate API call
function fetchData() {
    showLoading();
    setTimeout(() => {
        hideLoading();
        // Update UI with data
    }, 1500);
}
```

### Step 6: Test and Refine

**Testing checklist:**

**Visual Testing:**
- [ ] All colors match design system specifications
- [ ] Typography is consistent (sizes, weights, line heights)
- [ ] Spacing and alignment are correct
- [ ] Icons are properly sized and colored
- [ ] Images load correctly
- [ ] Shadows and borders are subtle and appropriate

**Interaction Testing:**
- [ ] All buttons are clickable and provide feedback
- [ ] Navigation works between pages/tabs
- [ ] Forms validate input correctly
- [ ] Modals open and close properly
- [ ] Dropdowns expand and collapse
- [ ] Hover states work on desktop
- [ ] Active/pressed states work on mobile

**Responsive Testing:**
- [ ] Layout adapts to different screen sizes
- [ ] Text remains readable at all sizes
- [ ] Touch targets are at least 44px × 44px on mobile
- [ ] No horizontal scrolling on mobile
- [ ] Content doesn't overflow containers

**Content Testing:**
- [ ] All text is readable and makes sense
- [ ] Data reflects realistic scenarios
- [ ] Empty states are shown when appropriate
- [ ] Error messages are clear and helpful
- [ ] Loading states provide feedback

**Accessibility Testing:**
- [ ] Color contrast meets WCAG AA standards (4.5:1 for text)
- [ ] Interactive elements are keyboard accessible
- [ ] Focus states are visible
- [ ] Alt text provided for images
- [ ] Semantic HTML used throughout

**Refinement actions:**
- Fix any visual inconsistencies
- Improve interaction feedback
- Optimize for performance (minimize CSS/JS)
- Add comments for complex logic
- Document any assumptions or limitations

### Step 7: Present and Iterate

**Present the prototype:**

1. **Save prototype files** to appropriate directory:
   ```
   docs/{date}-{feature-name}/prototypes/
   │  ├─ index.html
   │  ├─ styles.css (if separate)
   │  ├─ script.js (if separate)
   │  └─ README.md (usage instructions)
   ```

2. **Provide usage instructions:**
   ```markdown
   # [Feature Name] Prototype

   ## How to View

   1. Open `index.html` in a web browser
   2. For mobile view: Open browser DevTools (F12) and toggle device toolbar
   3. Recommended viewport: 375px × 812px (iPhone X)

   ## Features Demonstrated

   - [Feature 1]: [Description]
   - [Feature 2]: [Description]
   - [Feature 3]: [Description]

   ## Interactions

   - Click [element] to [action]
   - Navigate between pages using [navigation method]
   - Fill out forms and see validation

   ## Design System

   - Style: [WeChat Work / iOS Native / Material Design / Ant Design Mobile]
   - Primary Color: [#HEX]
   - Typography: [Font family]

   ## Notes

   - This is a prototype, not production code
   - Some interactions are simulated (no real backend)
   - Focus is on UI/UX validation, not functionality
   ```

3. **Highlight key design decisions:**
   - "I chose WeChat Work style because the PRD targets Chinese enterprise users"
   - "The dashboard prioritizes quick access to top 3 features based on user stories"
   - "Form validation follows the acceptance criteria in the PRD"

4. **Offer iteration options:**
   - "Would you like to adjust any colors or spacing?"
   - "Should I add more pages or interactions?"
   - "Do you want to see alternative layouts?"
   - "Any specific user flow you'd like me to refine?"

**Iterate based on feedback:**
- Update colors, typography, or spacing
- Add or remove features
- Refine interactions or animations
- Create alternative design variations
- Adjust for different screen sizes

## Design System Specifications

### WeChat Work Style

**Colors:**
- Primary: `#3478F6` (Tech Blue)
- Link: `#576B95` (Link Blue)
- Success: `#07C160` (Green)
- Warning: `#FF976A` (Orange)
- Error: `#FA5151` (Red)
- Background: `#F5F5F5` (Light Gray)
- Card: `#FFFFFF` (White)
- Text Primary: `#191919` (Near Black)
- Text Secondary: `#8C8C8C` (Gray)
- Border: `#E5E5E5` (Light Gray)

**Typography:**
- Font Family: `PingFang SC, -apple-system, BlinkMacSystemFont, sans-serif`
- Title: `18px, bold`
- Body: `15px, normal`
- Caption: `13px, normal`
- Line Height: `1.5`

**Components:**
- Navigation Bar: `44px` height
- List Item: `64px` height
- Button: `44px` height, `4px` border radius
- Card: `10px` border radius, subtle shadow
- Icon: `24px` or `48px` size

### iOS Native Style

**Colors:**
- System Blue: `#007AFF`
- System Green: `#34C759`
- System Red: `#FF3B30`
- System Gray: `#8E8E93`
- Background: `#F2F2F7`
- Card: `#FFFFFF`
- Text: `#000000`
- Secondary Text: `#3C3C43` (60% opacity)

**Typography:**
- Font Family: `SF Pro, -apple-system, sans-serif`
- Large Title: `34px, bold`
- Title: `28px, bold`
- Headline: `17px, semibold`
- Body: `17px, normal`
- Caption: `12px, normal`

**Components:**
- Navigation Bar: `44px` (compact) or `96px` (large title)
- List Row: `44px` minimum height
- Button: `50px` height, `10px` border radius
- Card: `10px` border radius
- Icon: `28px` SF Symbols

### Material Design Style

**Colors:**
- Primary: `#6200EE` (Purple)
- Secondary: `#03DAC6` (Teal)
- Error: `#B00020` (Red)
- Background: `#FFFFFF`
- Surface: `#FFFFFF`
- On Primary: `#FFFFFF`
- On Background: `#000000`

**Typography:**
- Font Family: `Roboto, sans-serif`
- H1: `96px, light`
- H2: `60px, light`
- H3: `48px, normal`
- H4: `34px, normal`
- H5: `24px, normal`
- H6: `20px, medium`
- Body 1: `16px, normal`
- Body 2: `14px, normal`

**Components:**
- App Bar: `56px` (mobile) or `64px` (desktop)
- List Item: `48px` or `72px` height
- Button: `36px` height, `4px` border radius
- Card: `4px` border radius, elevation
- FAB: `56px` diameter
- Icon: `24px` Material Icons

### Ant Design Mobile Style

**Colors:**
- Primary: `#108EE9` (Blue)
- Success: `#00A854` (Green)
- Warning: `#FFBF00` (Yellow)
- Error: `#F04134` (Red)
- Background: `#F5F5F9`
- Card: `#FFFFFF`
- Text: `#000000`
- Secondary Text: `#888888`
- Border: `#DDDDDD`

**Typography:**
- Font Family: `SF UI Display, -apple-system, sans-serif`
- Title: `18px, medium`
- Body: `14px, normal`
- Caption: `12px, normal`

**Components:**
- Navigation Bar: `45px` height
- List Item: `44px` height
- Button: `47px` height, `5px` border radius
- Card: `2px` border radius
- Icon: `22px` size

## Best Practices

**1. Start with Mobile-First:**
Design for mobile screens first, then enhance for larger screens. Most users will view prototypes on mobile devices.

**2. Use Real Content:**
Avoid "Lorem ipsum" and placeholder text. Use actual feature names, realistic data, and meaningful labels from the PRD.

**3. Show Multiple States:**
Include empty states, loading states, error states, and success states to demonstrate complete user flows.

**4. Keep It Simple:**
Focus on core functionality. Don't over-engineer interactions or add unnecessary features not in the PRD.

**5. Maintain Consistency:**
Use the same colors, typography, spacing, and component styles throughout the prototype.

**6. Optimize Performance:**
Minimize CSS and JavaScript. Use CDNs for libraries. Avoid large images or complex animations.

**7. Document Assumptions:**
If you make design decisions not specified in the PRD, document them clearly for stakeholder review.

**8. Test on Real Devices:**
Encourage users to test prototypes on actual mobile devices, not just desktop browsers.

**9. Iterate Quickly:**
Prototypes are meant to be disposable. Make changes quickly based on feedback without worrying about perfect code.

**10. Focus on UX Validation:**
The goal is to validate user experience and design decisions, not to write production-ready code.

## Common Patterns

### Pattern 1: Enterprise Dashboard (WeChat Work)

**Structure:**
```
├─ Top Navigation (44px)
│  ├─ Title (center)
│  └─ Menu Icon (right)
├─ Quick Access Grid (4 columns)
│  └─ Icon Buttons (48px icons)
├─ Data Summary Cards
│  └─ Metric Cards (horizontal layout)
├─ Feature List
│  └─ List Items (64px height, icon + text + arrow)
└─ Bottom Tab Bar (50px)
   └─ 5 Tabs (icon + label)
```

**Key Elements:**
- Tech blue (#3478F6) for primary actions
- White cards on light gray background
- Right arrow indicators for navigation
- Badge notifications on icons

### Pattern 2: iOS Content App

**Structure:**
```
├─ Large Title Navigation (96px when expanded)
│  └─ Title (34px bold)
├─ Search Bar (optional)
├─ Content Cards
│  ├─ Image (16:9 aspect ratio)
│  ├─ Title (17px semibold)
│  └─ Description (15px regular)
└─ Tab Bar (50px)
   └─ 4-5 Tabs (SF Symbols icons)
```

**Key Elements:**
- System blue (#007AFF) for interactive elements
- Generous whitespace (20px margins)
- Subtle dividers with left inset
- Pull-to-refresh gesture

### Pattern 3: Android Material App

**Structure:**
```
├─ Top App Bar (56px)
│  ├─ Menu Icon (left)
│  ├─ Title (center or left)
│  └─ Action Icons (right)
├─ Content Area
│  └─ Cards with Elevation
└─ FAB (Floating Action Button)
   └─ Primary Action (56px diameter)
```

**Key Elements:**
- Bold primary color (#6200EE)
- Elevation shadows (2dp, 4dp, 8dp)
- Ripple effects on tap
- 16dp grid system

### Pattern 4: Enterprise Form App (Ant Design Mobile)

**Structure:**
```
├─ Navigation Bar (45px)
│  ├─ Back Button (left)
│  └─ Title (center)
├─ Form Sections
│  ├─ Section Header
│  └─ Form Fields (grouped)
│     ├─ Label
│     ├─ Input
│     └─ Error Message
└─ Fixed Bottom Bar
   └─ Submit Button (47px height)
```

**Key Elements:**
- Professional blue (#108EE9)
- Dense information layout
- Clear field labels and validation
- Sticky bottom action bar

## Troubleshooting

**Issue: PRD lacks UI/UX details**
**Solution:** Use design system defaults, create reasonable assumptions, document decisions, and ask user for confirmation.

**Issue: Too many features to prototype**
**Solution:** Focus on MVP features first, prioritize primary user flows, create separate prototypes for different phases.

**Issue: Design system not specified**
**Solution:** Recommend based on platform, user demographics, and application type. Explain rationale and offer alternatives.

**Issue: Prototype doesn't match user expectations**
**Solution:** Gather specific feedback, identify misalignments with PRD, iterate quickly, offer alternative designs.

**Issue: Interactions too complex to implement**
**Solution:** Simplify to essential interactions, use static mockups for complex animations, focus on UX validation over technical implementation.

**Issue: Responsive design challenges**
**Solution:** Start with single target viewport (mobile or desktop), add responsive features incrementally, test on real devices.

## Resources

This skill includes reference materials and assets:

### references/
- `design-systems.md`: Detailed specifications for all supported design systems
- `component-library.md`: Reusable HTML/CSS component snippets
- `interaction-patterns.md`: Common JavaScript interaction patterns

### assets/
- `templates/`: Starter HTML templates for different page types
- `icons/`: Common icon sets (SVG)
- `images/`: Placeholder images for prototypes

---

**Note:** This skill generates actual prototype code (HTML/CSS/JavaScript), not just design specifications. The output is viewable in a web browser and suitable for stakeholder demos and UX validation.
