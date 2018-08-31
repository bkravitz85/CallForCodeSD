<aura:application >

    <ltng:require   scripts="/resource/carboncomponents/carboncomponents.js"
                  styles="/resource/carboncomponentscss/carboncomponentscss.css"
                    afterScriptsLoaded="{!c.afterScriptsLoaded}"/>
    <nav class="bx--cloud-header">
  <div class="bx--cloud-header__wrapper">
    <button class="bx--cloud-header__app-menu" type="button">

    </button>
    <a href="www.google.com" class="bx--cloud-header-brand">
      <div class="bx--cloud-header-brand__icon">

      </div>
      <h4 class="bx--cloud-header-brand__text">
        IBM
        <span>Cloud</span>
      </h4>
    </a>
    <ul class="bx--cloud-header-list">
      <li class="bx--cloud-header-list__item">
        <a class="bx--cloud-header-list__link" href="www.google.com">Catalog</a>
        <a class="bx--cloud-header-list__link" href="www.google.com">Support</a>
        <a class="bx--cloud-header-list__link" href="www.google.com">Manage</a>
      </li>
    </ul>
  </div>
  <div class="bx--cloud-header__wrapper">
    <ul class="bx--cloud-header-list">
      <li class="bx--cloud-header-list__item bx--cloud-header-list__item--icon">
        <button class="bx--cloud-header-list__btn" type="button">

        </button>
      </li>
      <li class="bx--cloud-header-list__item bx--cloud-header-list__item--icon">
        <button class="bx--cloud-header-list__btn" type="button">

        </button>
      </li>
      <li class="bx--cloud-header-list__item bx--cloud-header-list__item--icon">
        <button class="bx--cloud-header-list__btn" type="button">

        </button>
      </li>
      <li class="bx--cloud-header-list__item bx--cloud-header-list__item--icon">
        <button class="bx--cloud-header-list__btn" type="button">
        </button>
      </li>
    </ul>
  </div>
</nav>

 </aura:application>